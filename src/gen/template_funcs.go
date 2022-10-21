package main

import "github.com/gertd/go-pluralize"

var pl = pluralize.NewClient()

func Pluralize(w string) string {
	return pl.Plural(w)
}
