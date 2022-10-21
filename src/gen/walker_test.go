package main

import (
	"testing"
)

func TestSubstitution(t *testing.T) {
	variables := map[string]interface{}{
		// "MODULE_NAME":    "test_mod",
		"SOMETHING_ELSE": 42,
		// TODO: Verify if Unmarshal serializes to int
	}

	path := "./../some/folder/{{MODULE_NAME}}"

	result := SubstituteVariables(path, variables)

	if result != "./../some/folder/test_mod" {
		t.Errorf("unexpected result. got %s", result)
	}

}
