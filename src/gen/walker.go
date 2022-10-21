package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Walker struct {
}

func (w *Walker) Walk(inPath, outPath string, variables map[string]interface{}) {
	filepath.Walk(inPath, func(path string, info fs.FileInfo, err error) error {

		out := strings.ReplaceAll(path, inPath, outPath)
		fmt.Println(out)

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
