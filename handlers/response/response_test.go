package response

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustom(t *testing.T) {
	data := struct {
		UID string `json:"user"`
	}{
		UID: "5333888563",
	}
	expectedRespBody := &body{
		Data:   data,
		Status: http.StatusOK,
	}
	b, err := json.MarshalIndent(expectedRespBody, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	Custom(data, http.StatusOK, w)

	if !bytes.Equal(w.Body.Bytes(), b) {
		t.Errorf("expected respBody to be %q, got: %q", b, w.Body.Bytes())
	}
}

func TestOK(t *testing.T) {
	w := httptest.NewRecorder()

	OK(w)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code to be %d, got: %d", http.StatusOK, w.Code)
	}
}

func TestOptions(t *testing.T) {
	w := httptest.NewRecorder()

	optionMethods := "GET, POST, HEAD, OPTIONS"
	Options(optionMethods, w)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code to be %d, got: %d", http.StatusOK, w.Code)
	}

	accessControlMethods := w.Header()["Access-Control-Allow-Methods"][0]
	if accessControlMethods != optionMethods {
		t.Errorf("expected access control methods to be %q, got: %q", optionMethods, accessControlMethods)
	}
	allowMethods := w.Header()["Allow"][0]
	if allowMethods != optionMethods {
		t.Errorf("expected allow methods to be %q, got: %q", optionMethods, allowMethods)
	}
}
