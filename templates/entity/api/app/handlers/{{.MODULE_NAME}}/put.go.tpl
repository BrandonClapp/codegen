package {{ .MODULE_TYPE | ToLower }}

import (
	"net/http"

	"github.com/brandonclapp/budget/core/data"
	coreHttp "github.com/brandonclapp/budget/core/http"
	"github.com/gorilla/mux"
)

func putOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		msg := "id must be specified"
		coreHttp.WriteJsonResponse(w, &msg, http.StatusNotFound)
		return
	}

	var entity = &{{.MODULE_TYPE}}{}
	coreHttp.ParseBody(w, r, entity)

	updates := map[string]interface{}{
		"name":  entity.Name,
		"color": entity.Color,
	}

	updated, err := {{.MODULE_TYPE | Pluralize}}.UpdateOne(updates, data.IDEquals(id))

	if err != nil {
		coreHttp.WriteJsonResponse(w, &err, http.StatusInternalServerError)
		return
	}

	coreHttp.WriteJsonResponse(w, updated, http.StatusOK)
}
