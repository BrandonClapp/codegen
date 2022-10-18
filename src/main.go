package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed template/*
var template embed.FS

func main() {

	// TODO: Check for the presense of a gen.config.json file

	// Generate("person", "first, last, age")
	fs.WalkDir(template, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// buf, err := ioutil.ReadFile(path)
		buf, err := fs.ReadFile(template, path)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Println(string(buf))

		// fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		return nil
	})
}
