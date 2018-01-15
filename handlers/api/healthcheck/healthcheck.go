package healthcheck

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewrk/design-brain-api/handlers/config"
	"github.com/nicewrk/design-brain-api/handlers/response"
)

// Handler handles GET, HEAD, and OPTIONS requests to /healthcheck.
func Handler(cfg *config.Config) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// Switch on HTTP method.
		switch req.Method {

		// GET
		case http.MethodGet:
			response.Custom(nil, http.StatusOK, w)

			// HEAD
		case http.MethodHead:
			response.OK(w)

			// OPTIONS
		case http.MethodOptions:
			response.Options(optionMethods, w)
		}
	}
}
