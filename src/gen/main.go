package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Template struct {
	Name      string
	InDir     string
	OutDir    string
	Variables interface{}
}

type Config struct {
	Templates []Template
}

func main() {

	args := getArgs()
	if len(args) == 1 {
		panic("Template argument must be specified")
	}

	tpl := args[1]

	// Read in config from current directory
	buf, err := ioutil.ReadFile("gen.config.json")
	if err != nil {
		panic("gen.config.json file not in this directory")
	}

	// Unmarshal to config type
	config := Config{}
	err = json.Unmarshal(buf, &config)
	if err != nil {
		panic("Unable to serialize gen.config.json")
	}

	template := getTemplate(&config.Templates, tpl)

	if template == nil {
		panic(fmt.Sprintf("Unable to find template with name %s in gen.config.json", tpl))
	}

	configMap := template.Variables.(map[string]interface{})

	fmt.Println(configMap["MODULE_NAME"])

}

func getTemplate(templates *[]Template, name string) *Template {
	for _, tpl := range *templates {
		if tpl.Name == name {
			return &tpl
		}
	}
	return nil
}

func getArgs() []string {
	args := []string{}
	for _, v := range os.Args {
		// filter out debug_bin argument when debugging
		if strings.Contains(v, "debug_bin") {
			continue
		}
		args = append(args, v)
	}

	return args
}
