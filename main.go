package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nicewrk/design-brain-api/dotenv"
	"github.com/nicewrk/design-brain-api/handlers"
	"github.com/nicewrk/design-brain-api/handlers/config"
	"github.com/nicewrk/design-brain-api/newrelic"
	"github.com/nicewrk/design-brain-api/store"
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

	// Initialize store client.
	storeClient, err := store.NewClient()
	if err != nil {
		log.Fatalf("error: initializing store client: %s.", err)
	}

	// Initialize handler configuration.
	cfg := &config.Config{
		NewRelicApp: app,
		StoreClient: storeClient,
	}

	// Initialize router.
	router := handlers.NewRouter(cfg)

	// Start server.
	log.Fatal(http.ListenAndServe(":8080", router))
}
