package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nicewrk/design-brain-api/handlers/api/healthcheck"
	"github.com/nicewrk/design-brain-api/handlers/api/users"
	"github.com/nicewrk/design-brain-api/handlers/config"
)

// NewRouter configures and returns a new *httprouter.Router.
func NewRouter(cfg *config.Config) *httprouter.Router {
	router := httprouter.New()

	// /healthcheck
	router.GET("/healthcheck", healthcheck.Handler(cfg))
	router.HEAD("/healthcheck", healthcheck.Handler(cfg))
	router.OPTIONS("/healthcheck", healthcheck.Handler(cfg))

	// /users
	router.GET("/users/:username", users.Handler(cfg))
	router.POST("/users", users.Handler(cfg))
	router.HEAD("/users", users.Handler(cfg))
	router.OPTIONS("/users", users.Handler(cfg))

	return router
}
