package users

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewrk/design-brain-api/handlers/config"
	"github.com/nicewrk/design-brain-api/handlers/response"
	"github.com/nicewrk/design-brain-api/store/models"
)

// Handler handles GET, POST, HEAD, and OPTIONS requests to /users.
func Handler(cfg *config.Config) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// Switch on HTTP method.
		switch req.Method {

		// GET
		case http.MethodGet:
			u, err := handle(cfg, w, req, ps)
			if err != nil {
				response.Bad(err, errorStatusCode(err), w)
			} else {
				response.Custom(u, http.StatusOK, w)
			}

		// POST
		case http.MethodPost:
			u, err := handle(cfg, w, req, ps)
			if err != nil {
				response.Bad(err, errorStatusCode(err), w)
			} else {
				response.Custom(u, http.StatusCreated, w)
			}

			// HEAD
		case http.MethodHead:
			response.OK(w)

			// OPTIONS
		case http.MethodOptions:
			response.Options(optionMethods, w)
		}
	}
}

func handle(cfg *config.Config, w http.ResponseWriter, req *http.Request, ps httprouter.Params) (*models.User, error) {
	// Switch on HTTP method.
	switch req.Method {

	// GET
	case http.MethodGet:
		u := &models.User{}
		u.Username = ps.ByName("username")
		u, err := u.Select(cfg.StoreClient)
		if err != nil {
			return nil, errUserNotFound
		}
		return u, nil

	// POST
	case http.MethodPost:
		u := &models.User{}
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			return nil, errUnableToProcessUserPayload
		}
		err = u.Insert(cfg.StoreClient)
		if err != nil {
			return nil, storeErrors(err)
		}
		u, err = u.Select(cfg.StoreClient)
		if err != nil {
			return nil, storeErrors(err)
		}
		return u, nil
	}
	return nil, errUnknownProcessing
}
