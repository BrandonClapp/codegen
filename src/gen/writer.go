package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteTemplateOutput(outDir string, output *TemplateOutput) {
	// for each key, write to output
	for k, v := range *output {
		fileOutput := fmt.Sprintf("%s%s", outDir, k)
		if _, err := os.Stat(k); os.IsNotExist(err) {
			dir := filepath.Dir(fileOutput)

			// ensure directory exists
			os.MkdirAll(dir, os.ModePerm)
		}

		// write file to output destination
		err := ioutil.WriteFile(fileOutput, []byte(v), os.ModePerm)

		if err != nil {
			panic(err)
		}
	}

}
