package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewrk/design-brain-api/dotenv"
	"github.com/nicewrk/design-brain-api/handlers/healthcheck"
	"github.com/nicewrk/design-brain-api/newrelic"
)

func init() {
	err := dotenv.Run(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize New Relic.
	app, err := newrelic.Init(os.Getenv("APP_NAME"))
	if err != nil {
		log.Fatalf("error: initializing New Relic: %s.", err)
	}

	// Initialize router.
	router := httprouter.New()

	// /healthcheck
	router.GET(healthcheck.Path, app.WrapHandler(healthcheck.Handle))
	router.HEAD(healthcheck.Path, app.WrapHandler(healthcheck.Handle))
	router.OPTIONS(healthcheck.Path, app.WrapHandler(healthcheck.Handle))

	log.Fatal(http.ListenAndServe(":8080", router))
}
