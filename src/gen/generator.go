package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

type Vars map[string]interface{}

func Generate(templateRootPath, outputPath string, variables map[string]interface{}) error {

	// create temp directory
	if !pathExists(outputPath) {
		err := os.MkdirAll(outputPath, os.ModePerm)

		if err != nil {
			return err
		}
	}

	// iterate over every file and folder in the template directory
	err := filepath.Walk(templateRootPath, func(path string, info fs.FileInfo, err error) error {

		// substitute any variables found in files and folder paths
		resolvedPath, err := RenderTemplate(path, path, variables)

		if err != nil {
			return err
		}

		tempOut := strings.ReplaceAll(resolvedPath, templateRootPath, outputPath)

		// create directories in the output
		if info.IsDir() && !pathExists(tempOut) {
			fmt.Printf("Making dir at %s \n", tempOut)
			os.Mkdir(tempOut, os.ModePerm)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		// read the template file contents
		buf, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Printf("could not read file %s\n", path)
		}

		// transform the template file with variables
		content, err := RenderTemplate(path, string(buf), variables)

		if err != nil {
			// an error occured while transforming the template, roll back
			fmt.Printf("\nError: %s\n\nRemoving %s", err.Error(), outputPath)
			cleanup(outputPath)
			return err
		}

		// remove .tpl extension from template
		tempOut = strings.ReplaceAll(tempOut, ".tpl", "")
		abs, _ := filepath.Abs(tempOut)

		// write the transformed file to the temp directory
		err = ioutil.WriteFile(abs, []byte(content), os.ModePerm)

		if err != nil {
			// fmt.Println(err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil

}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func cleanup(outputPath string) {
	err := os.RemoveAll(outputPath)

	if err != nil {
		panic(err)
	}
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
