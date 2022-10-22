package main

import "github.com/gertd/go-pluralize"

var pl = pluralize.NewClient()

func Pluralize(w string) string {
	return pl.Plural(w)
}

func IsLast(index int, len int) bool {
	return index+1 == len
}

func ToPostgresType(typ string) string {
	switch typ {
	case "string":
		return "text"
	case "int":
		return "integer"
	}

	return ""
}
