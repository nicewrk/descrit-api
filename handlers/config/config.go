package config

import (
	"github.com/nicewrk/design-brain-api/newrelic"
	"github.com/nicewrk/design-brain-api/store"
)

// Config holds the handler configurations.
type Config struct {
	NewRelicApp newrelic.Application
	StoreClient store.API
}
