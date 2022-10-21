package {{ .MODULE_TYPE | ToLower }}

import (
	"net/http"

	"github.com/brandonclapp/budget/core/data"
	coreHttp "github.com/brandonclapp/budget/core/http"
	"github.com/gorilla/mux"
)

// Get all
func get(w http.ResponseWriter, r *http.Request) {
	entities := {{.MODULE_TYPE | Pluralize }}.Get(nil)
	coreHttp.WriteJsonResponse(w, &entities, http.StatusOK)
}

// Get single
func getOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		msg := "id must be specified"
		coreHttp.WriteJsonResponse(w, &msg, http.StatusNotFound)
		return
	}

	{{.MODULE_NAME}}, err := {{.MODULE_TYPE | Pluralize}}.GetOne(data.IDEquals(id))

	if err != nil {
		coreHttp.WriteJsonResponse(w, &err, http.StatusNotFound)
		return
	}

	coreHttp.WriteJsonResponse(w, &{{.MODULE_NAME}}, http.StatusOK)
}
