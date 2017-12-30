package healthcheck

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	// Path is the path for simple healthchecks.
	Path = "/healthcheck"
)

// Handle handles GET, HEAD, and OPTIONS requests to /healthcheck.
func Handle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// Switch on HTTP method.
	switch req.Method {

	// GET
	case http.MethodGet:
		custom(nil, http.StatusOK, w)

		// HEAD
	case http.MethodHead:
		ok(w)

		// OPTIONS
	case http.MethodOptions:
		options(optionMethods, w)
	}
}
