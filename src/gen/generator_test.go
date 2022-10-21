package main

import (
	"fmt"
	"testing"
)

func TestSubstitution(t *testing.T) {
	variables := map[string]interface{}{
		"MODULE_NAME":    "person",
		"SOMETHING_ELSE": 42,
		// TODO: Verify if Unmarshal serializes to int
	}

	path := "./../some/folder/{{.MODULE_NAME}}"

	result, err := RenderTemplate(path, path, variables)

	if err != nil {
		t.Error(err)
	}

	if result != "./../some/folder/person" {
		t.Errorf("unexpected result. got %s", result)
	}
}

func TestTemplateParse(t *testing.T) {
	variables := map[string]interface{}{
		"MODULE_TYPE": "Person",
	}

	tpl := `
package {{ .MODULE_TYPE | ToLower }}

import (
"net/http"

	"github.com/brandonclapp/budget/core/data"
	coreHttp "github.com/brandonclapp/budget/core/http"
	"github.com/gorilla/mux"
)

// Get all
func get(w http.ResponseWriter, r *http.Request) {
	entities := {{.MODULE_TYPE | Pluralize}}.Get(nil)
	coreHttp.WriteJsonResponse(w, &entities, http.StatusOK)
}
`

	output, err := RenderTemplate("get", tpl, variables)

	if err != nil {
		t.Error("error parsing template")
	}

	fmt.Println(output)

}
