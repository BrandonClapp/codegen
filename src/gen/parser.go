package main

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

type Parser struct {
}

func (p *Parser) Parse(name, content string, variables Vars) (string, error) {
	t, err := template.New(name).Parse(content)
	if err != nil {
		err = errors.New(fmt.Sprintf("failure parsing template %s. \n error: %s", name, err.Error()))
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, variables)

	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
