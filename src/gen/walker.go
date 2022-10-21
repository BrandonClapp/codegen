package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Vars map[string]interface{}

type Walker struct {
}

func (w *Walker) Walk(inPath, outPath string, variables map[string]interface{}) {
	filepath.Walk(inPath, func(path string, info fs.FileInfo, err error) error {

		// substitute any variables found in files and folder paths
		path = SubstituteVariables(path, variables)

		// find the equivelent output path for this file or folder
		out := strings.ReplaceAll(path, inPath, outPath)

		// create directories in the output
		if info.IsDir() && !pathExists(out) {
			fmt.Printf("Making dir at %s", out)
			os.Mkdir(out, os.ModePerm)
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)

		// buf, err := ioutil.ReadFile(path)

		// if err != nil {
		// 	fmt.Printf("could not read file %s\n", path)
		// }

		return nil
	})
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func SubstituteVariables(path string, variables Vars) string {
	p := path
	for k, v := range variables {
		switch val := v.(type) {
		case string:
			// only string variables can be substituted
			p = strings.ReplaceAll(p, "{{"+k+"}}", val)
		default:
			continue
		}

	}

	return p
}
