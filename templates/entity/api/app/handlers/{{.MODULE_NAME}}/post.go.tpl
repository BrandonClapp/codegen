package {{ .MODULE_TYPE | ToLower }}

import (
	"log"
	"net/http"

	coreHttp "github.com/brandonclapp/budget/core/http"
)

func post(w http.ResponseWriter, r *http.Request) {
	var entity = &{{.MODULE_TYPE}}{}

	coreHttp.ParseBody(w, r, entity)

	created, err := {{.MODULE_TYPE | Pluralize}}.Create(entity)

	if err != nil {
		log.Fatalf("%s", err)
		coreHttp.WriteJsonResponse(w, &err, http.StatusInternalServerError)
		return
	}

	coreHttp.WriteJsonResponse(w, created, http.StatusOK)
}
