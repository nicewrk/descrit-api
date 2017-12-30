package newrelic

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"net/http"

	"github.com/julienschmidt/httprouter"
	newrelicagent "github.com/newrelic/go-agent"
)

// Application wraps newrelicagent.Application.
type Application struct {
	newrelicagent.Application
}

// Init configures a New Relic agent and then initializes and returns a New Relic application.
func Init(baseAppName string) (Application, error) {
	config := newrelicagent.NewConfig(appName(baseAppName), os.Getenv("NEWRELIC_LICENSE_KEY"))

	isProductionOrStaging, err := regexp.MatchString("(production|staging)", os.Getenv("ENVIRONMENT"))
	if err != nil {
		log.Fatalf("error: identifying environment: %s.", err)
	}
	config.Enabled = isProductionOrStaging

	app, err := newrelicagent.NewApplication(config)

	return Application{app}, err
}

func appName(baseAppName string) string {
	switch os.Getenv("ENVIRONMENT") {
	case "production":
		return baseAppName
	case "staging":
		return fmt.Sprintf("%s (staging)", baseAppName)
	}
	return fmt.Sprintf("%s (dev)", baseAppName)
}

// WrapHandler is middleware that wraps a httprouter.Handle function. It starts a NewRelic transaction
// and gracefully ends it. WrapHandler takes a httprouter.Handle as an argument and returns a httprouter.Handle.
func (app Application) WrapHandler(handle httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

		// Begin New Relic transaction.
		txn := app.StartTransaction(req.URL.RequestURI(), w, req)
		defer func() {
			err := txn.End()
			if err != nil {
				log.Printf("error: ending New Relic transaction | %s", err)
			}
		}()

		handle(w, req, ps)
	}
}
