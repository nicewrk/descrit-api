package newrelic

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestAppName(t *testing.T) {
	t.Parallel()

	baseAppName := "TA API"

	testCases := []struct {
		env  string
		want string
	}{
		{"", fmt.Sprintf("%s (dev)", baseAppName)},
		{"staging", fmt.Sprintf("%s (staging)", baseAppName)},
		{"production", baseAppName},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("appName in %s", tc.env), func(t *testing.T) {
			err := os.Setenv("ENVIRONMENT", tc.env)
			if err != nil {
				t.Fatal(fmt.Sprintf("error: setting ENVIRONMENT | %s", err))
			}
			defer func() {
				err := os.Unsetenv("ENVIRONMENT")
				if err != nil {
					t.Fatal(fmt.Sprintf("error: unsetting ENVIRONMENT | %s", err))
				}
			}()

			appName := appName(baseAppName)
			if appName != tc.want {
				t.Errorf("got %s; want %s", appName, tc.want)
			}
		})
	}
}

func TestWrapHandler(t *testing.T) {
	t.Parallel()

	app, err := Init("test app")
	if err != nil {
		t.Error(err)
	}

	// A simple test just to make sure calling doesn't blow up!
	app.WrapHandler(func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) { return })
}
