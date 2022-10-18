package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// nova gen category --queryable=Categories --columns=name,color
// Read all files in template

// Make replacements
// {{MODULE_NAME}} -> category
// {{MODULE_TYPE}} -> Category
// {{MODULE_TYPE_PLURAL}} -> Categories
// {{MODULE_COLUMNS}} -> "name", "color"
// {{MODULE_STRUCT_PROPS}} ->
// Name string `json:"name"`
// Color string `json:"color"`
// {{MODULE_SQL_PROP_MAPPINGS}} -> Change name, color into e.Name, e.Color
// {{MODULE_COLUMNS_SQL}} -> Change name, color into:
// name text
// color text

// Output to directory in app

func Generate(moduleName string, columns string) {

	variables := getVariables(moduleName, columns)
	outDir := fmt.Sprintf("../api/app/handlers/%s", moduleName)

	os.Mkdir(outDir, os.ModePerm)
	filepath.Walk("./template", func(path string, info fs.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		buf, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Printf("could not read file %s\n", path)
		}

		content := string(buf)

		content = strings.ReplaceAll(content, "{{MODULE_NAME}}", variables.ModuleName)
		content = strings.ReplaceAll(content, "{{MODULE_TYPE}}", variables.ModuleType)
		content = strings.ReplaceAll(content, "{{MODULE_TYPE_PLURAL}}", variables.ModuleTypePlural)
		content = strings.ReplaceAll(content, "{{MODULE_COLUMNS}}", variables.ModuleColumns)
		content = strings.ReplaceAll(content, "{{MODULE_STRUCT_PROPS}}", variables.ModuleStructProps)
		content = strings.ReplaceAll(content, "{{MODULE_SQL_PROP_MAPPINGS}}", variables.ModuleSqlPropMappings)
		content = strings.ReplaceAll(content, "{{MODULE_COLUMNS_SQL}}", variables.ModuleColumnsSql)

		fmt.Println(content)

		fn := strings.ReplaceAll(info.Name(), ".mustache", "")

		if fn == "module.go" {
			fn = moduleName + ".go"
		}

		out := fmt.Sprintf("%s/%s", outDir, fn)
		abs, _ := filepath.Abs(out)
		err = ioutil.WriteFile(abs, []byte(content), 0644)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		return nil
	})

}

type Variables struct {
	ModuleName            string
	ModuleType            string
	ModuleTypePlural      string
	ModuleSqlPropMappings string
	ModuleColumns         string
	ModuleStructProps     string
	ModuleColumnsSql      string
}

func getVariables(moduleName string, columns string) Variables {
	MODULE_NAME := moduleName
	MODULE_TYPE := strings.Title(MODULE_NAME)
	MODULE_TYPE_PLURAL := MODULE_TYPE + "s"

	// Split columns
	cols := strings.Split(strings.ReplaceAll(columns, " ", ""), ",")

	MODULE_SQL_PROP_MAPPINGS := genSqlPropMappings(cols)
	MODULE_COLUMNS := genColumns(cols)
	MODULE_STRUCT_PROPS := genModuleStructProps(cols)
	MODULE_COLUMNS_SQL := genModuleColumnsSql(cols)

	return Variables{
		ModuleName:            MODULE_NAME,
		ModuleType:            MODULE_TYPE,
		ModuleTypePlural:      MODULE_TYPE_PLURAL,
		ModuleSqlPropMappings: MODULE_SQL_PROP_MAPPINGS,
		ModuleColumns:         MODULE_COLUMNS,
		ModuleStructProps:     MODULE_STRUCT_PROPS,
		ModuleColumnsSql:      MODULE_COLUMNS_SQL,
	}
}

// Generates a string used in the Insert and Scan functions
// i.e.
// e.Name, e.Color
func genSqlPropMappings(columns []string) string {
	var set []string
	for _, v := range columns {
		prop := fmt.Sprintf(`e.%s`, strings.Title(v))
		set = append(set, prop)
	}
	colMap := strings.Join(set, ", ")
	return colMap
}

func genColumns(columns []string) string {
	var set []string

	for _, v := range columns {
		col := fmt.Sprintf(`"%s"`, v)
		set = append(set, col)
	}

	cols := strings.Join(set, ", ")
	return cols
}

func genModuleColumnsSql(columns []string) string {
	var set []string
	for _, v := range columns {
		col := fmt.Sprintf("%s text", v)
		set = append(set, col)
	}

	cols := strings.Join(set, ",\n")
	return cols
}

func genModuleStructProps(columns []string) string {
	var set []string
	for _, v := range columns {
		name := strings.Title(v)
		json := strings.ToLower(v)
		prop := fmt.Sprintf("%s string `json:\"%s\"`", name, json)
		set = append(set, prop)
	}

	props := strings.Join(set, "\n")
	return props
}
