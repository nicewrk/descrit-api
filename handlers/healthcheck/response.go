package healthcheck

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	optionMethods = "GET, HEAD, OPTIONS"
)

// body respresents the body of the JSON response sent to clients.
type body struct {
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
	Status int         `json:"status,omitempty"`
}

// meta represents the self-referential component of a specific portion of a JSON response.
type meta struct {
	Description string `json:"description,omitempty"`
}

// custom writes a custom JSON response back to the client.
func custom(data interface{}, statusCode int, w http.ResponseWriter) {
	respBody := &body{
		Data:   data,
		Status: statusCode,
	}
	write(respBody, w)
}

// ok writes a simple 200 HEAD response back to the client.
func ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// options writes a simple 200 OPTIONS response back to the client.
func options(allowedMethods string, w http.ResponseWriter) {
	w.Header().Set("Allow", allowedMethods)
	w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
	w.WriteHeader(http.StatusOK)
}

func write(respBody *body, w http.ResponseWriter) {
	b, err := json.MarshalIndent(respBody, "", "\t")
	if err != nil {
		log.Printf("error: marshalling respBody: %s", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(respBody.Status)
		_, err = w.Write(b)
		if err != nil {
			log.Printf("error: writing respBody: %s", err)
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
	}
}
