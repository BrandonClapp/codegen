package main

import "github.com/gertd/go-pluralize"

var pl = pluralize.NewClient()

func Pluralize(w string) string {
	return pl.Plural(w)
}

func IsLast(index int, len int) bool {
	return index+1 == len
}

// TODO: Don't do this. Find a better way to support this directly from template,
// rather than baking this logic into the binary
// Usage: {{ "string" | ToPostgresType }} yields "text"
func ToPostgresType(typ string) string {
	switch typ {
	case "string":
		return "text"
	case "int":
		return "integer"
	}

	return ""
}
