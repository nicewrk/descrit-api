package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// body respresents the body of the JSON response sent to clients.
type body struct {
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
	Status int         `json:"status,omitempty"`
}

// Bad writes a bad JSON response back to the client.
func Bad(err error, statusCode int, w http.ResponseWriter) {
	respBody := &body{
		Error:  err.Error(),
		Status: statusCode,
	}
	log.Println(err)
	write(respBody, w)
}

// Custom writes a custom JSON response back to the client.
func Custom(data interface{}, statusCode int, w http.ResponseWriter) {
	respBody := &body{
		Data:   data,
		Status: statusCode,
	}
	write(respBody, w)
}

// OK writes a simple 200 HEAD response back to the client.
func OK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// Options writes a simple 200 OPTIONS response back to the client.
func Options(allowedMethods string, w http.ResponseWriter) {
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
