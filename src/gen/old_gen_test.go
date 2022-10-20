package main

import "testing"

func TestGen(t *testing.T) {
	Generate("person", "first, last, weight, height, age, color")
}
