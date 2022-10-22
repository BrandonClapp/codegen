package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

type Vars map[string]interface{}

type TemplateOutput map[string]string

func Generate(templateRootPath string, variables map[string]interface{}) (*TemplateOutput, error) {

	InMemoryFS := make(TemplateOutput)

	// iterate over every file and folder in the template directory
	err := filepath.Walk(templateRootPath, func(path string, info fs.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		// substitute any variables found in files and folder paths
		resolvedPath, err := RenderTemplate(path, path, variables)

		if err != nil {
			return err
		}

		resolvedPath = strings.ReplaceAll(resolvedPath, templateRootPath, "")

		// read the template file contents
		buf, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Printf("could not read file %s\n", path)
		}

		// transform the template file with variables
		content, err := RenderTemplate(path, string(buf), variables)

		if err != nil {
			// an error occured while transforming the template, abort
			fmt.Printf("\nError tranforming %s:\n%s\n", path, err.Error())
			return err
		}

		// remove .tpl extension from template
		resolvedPath = strings.ReplaceAll(resolvedPath, ".tpl", "")

		// add the transformed file to the map
		InMemoryFS[resolvedPath] = string(content)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &InMemoryFS, nil

}

// https://github.com/iancoleman/strcase
var funcMap = template.FuncMap{
	"ToLower":          strings.ToLower,          // anykindofstring
	"ToKebab":          strcase.ToKebab,          // any-kind-of-string
	"ToScreamingKebab": strcase.ToScreamingKebab, // ANY-KIND-OF-STRING
	"ToSnake":          strcase.ToSnake,          // any_kind_of_string
	"ToScreamingSnake": strcase.ToScreamingSnake, // ANY_KIND_OF_STRING
	"ToCamel":          strcase.ToCamel,          // AnyKindOfString
	"ToLowerCamel":     strcase.ToLowerCamel,     // anyKindOfString
	"Pluralize":        Pluralize,                // Person -> People
	"IsLast":           IsLast,
	"ToPostgresType":   ToPostgresType, // Convert to postgres column type
}

func RenderTemplate(name, content string, variables Vars) (string, error) {
	t := template.Must(template.New(name).
		Funcs(funcMap).             // provide helper functions to templates
		Option("missingkey=error"). // error if template has missing variables
		Parse(content))

	var tpl bytes.Buffer
	err := t.Execute(&tpl, variables)

	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
