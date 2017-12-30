package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewrk/design-brain-api/dotenv"
	"github.com/udacity/ta-api/handlers/healthcheck"
)

func init() {
	err := dotenv.Run(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize router.
	router := httprouter.New()

	// /healthcheck
	router.GET(healthcheck.Path, app.WrapHandler(healthcheck.Handle))
	router.HEAD(healthcheck.Path, app.WrapHandler(healthcheck.Handle))
	router.OPTIONS(healthcheck.Path, app.WrapHandler(healthcheck.Handle))

	log.Fatal(http.ListenAndServe(":8080", router))
}
