package {{MODULE_NAME}}

import (
	"net/http"

	"github.com/brandonclapp/budget/core/data"
	coreHttp "github.com/brandonclapp/budget/core/http"
	"github.com/gorilla/mux"
)

func deleteOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		msg := "id must be specified"
		coreHttp.WriteJsonResponse(w, &msg, http.StatusNotFound)
		return
	}

	deletedID, err := {{MODULE_TYPE_PLURAL}}.DeleteOne(data.IDEquals(id))

	if err != nil {
		coreHttp.WriteJsonResponse(w, &err, http.StatusNotFound)
		return
	}

	coreHttp.WriteJsonResponse(w, &deletedID, http.StatusOK)
}
