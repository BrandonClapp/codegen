package main

import (
	"fmt"
	"testing"
)

func TestTemplateParse(t *testing.T) {
	variables := map[string]interface{}{
		"MODULE_NAME":        "person",
		"MODULE_TYPE_PLURAL": "People",
	}

	tpl := `
package {{.MODULE_NAME}}

import (
"net/http"

	"github.com/brandonclapp/budget/core/data"
	coreHttp "github.com/brandonclapp/budget/core/http"
	"github.com/gorilla/mux"
)

// Get all
func get(w http.ResponseWriter, r *http.Request) {
	entities := {{.MODULE_TYPE_PLURAL}}.Get(nil)
	coreHttp.WriteJsonResponse(w, &entities, http.StatusOK)
}
`

	p := &Parser{}
	output, err := p.Parse("get", tpl, variables)

	if err != nil {
		t.Error("error parsing template")
	}

	fmt.Println(output)

}
