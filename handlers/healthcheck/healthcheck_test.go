package healthcheck

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHealthCheckHandler(t *testing.T) {
	t.Parallel()

	// Initialize router.
	router := httprouter.New()

	testCases := []struct {
		method      string
		status      int
		contentType string
		respBody    *body
	}{
		{http.MethodGet, http.StatusOK, "application/json", &body{Status: 200}},
		{http.MethodHead, http.StatusOK, "", nil},
		{http.MethodOptions, http.StatusOK, "", nil},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(tc.method, Path, nil)
		w := httptest.NewRecorder()

		router.Handle(tc.method, Path, Handle)
		router.ServeHTTP(w, req)

		resp := w.Result()

		t.Run(fmt.Sprintf("%s %s", tc.method, Path), func(t *testing.T) {
			if resp.StatusCode != tc.status {
				t.Errorf("expected %d, got: %d", tc.status, resp.StatusCode)
			}
			contentType := resp.Header.Get("Content-Type")
			if contentType != tc.contentType {
				t.Errorf("expected %s, got: %s", tc.contentType, contentType)
			}

			switch tc.method {
			case http.MethodGet:
				respBody := &body{}
				err := json.NewDecoder(resp.Body).Decode(&respBody)
				if err != nil {
					t.Fatal(err)
				}

				if respBody.Status != tc.respBody.Status {
					t.Errorf("expected %d, got: %d", tc.respBody.Status, respBody.Status)
				}
			case http.MethodOptions:
				accessControlMethods := resp.Header.Get("Access-Control-Allow-Methods")
				if accessControlMethods != optionMethods {
					t.Errorf("expected %s, got: %s", optionMethods, accessControlMethods)
				}
				allowMethods := resp.Header.Get("Allow")
				if allowMethods != optionMethods {
					t.Errorf("expected %s, got: %s", allowMethods, optionMethods)
				}
			}

		})
	}
}
