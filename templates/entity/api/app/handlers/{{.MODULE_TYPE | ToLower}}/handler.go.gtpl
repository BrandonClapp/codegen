package {{ .MODULE_TYPE | ToLower }}

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Get all resources
	if r.Method == "GET" {
		get(w, r)
	}

	// Create new resource
	if r.Method == "POST" {
		post(w, r)
	}
}

func IndividualHandler(w http.ResponseWriter, r *http.Request) {
	// Get single resource
	if r.Method == "GET" {
		getOne(w, r)
	}

	// Update single resource
	if r.Method == "PUT" {
		putOne(w, r)
	}

	// Delete single resource
	if r.Method == "DELETE" {
		deleteOne(w, r)
	}
}
