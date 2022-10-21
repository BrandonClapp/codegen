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
)

type Vars map[string]interface{}

type Walker struct {
}

var tempDir string = "./.codegen/tmp"

func withTempDirectory(fn func(tmpDir string) error) error {
	if !pathExists(tempDir) {
		// create temp directory
		err := os.MkdirAll(tempDir, os.ModePerm)

		if err != nil {
			return err
		}
	}

	err := fn(tempDir)

	if err != nil {
		return err
	}

	cleanup()

	return nil
}

func (w *Walker) Walk(inPath, outPath string, variables map[string]interface{}) error {

	withTempDirectory(func(tmpDir string) error {
		err := filepath.Walk(inPath, func(path string, info fs.FileInfo, err error) error {

			// substitute any variables found in files and folder paths
			resolvedPath, err := RenderTemplate(path, path, variables)

			if err != nil {
				return err
			}

			tempOut := strings.ReplaceAll(resolvedPath, inPath, tmpDir)

			// create directories in the output
			if info.IsDir() && !pathExists(tempOut) {
				fmt.Printf("Making dir at %s \n", tempOut)
				os.Mkdir(tempOut, os.ModePerm)
				return nil
			}

			if info.IsDir() {
				return nil
			}

			// found a file
			buf, err := ioutil.ReadFile(path)

			if err != nil {
				fmt.Printf("could not read file %s\n", path)
			}

			content, err := RenderTemplate(path, string(buf), variables)

			if err != nil {
				fmt.Printf("\nError: %s\n\nRemoving %s", err.Error(), tmpDir)
				cleanup()
				return err
			}

			// remove .tpl extension from template
			tempOut = strings.ReplaceAll(tempOut, ".tpl", "")
			abs, _ := filepath.Abs(tempOut)
			err = ioutil.WriteFile(abs, []byte(content), 0644)

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
	})

	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func cleanup() {
	err := os.RemoveAll(tempDir)

	if err != nil {
		panic(err)
	}
}

func RenderTemplate(name, content string, variables Vars) (string, error) {

	// Map name formatDate to formatDate function above
	// var funcMap = template.FuncMap{
	// 	"formatDate": formatDate,
	// }

	// t = template.Must(template.New("template-07.txt").Funcs(funcMap).ParseFiles("template-08.txt"))

	t := template.Must(template.New(name).Option("missingkey=error").Parse(content))

	var tpl bytes.Buffer
	err := t.Execute(&tpl, variables)

	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
