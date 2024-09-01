package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpErrors string

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "applicatin/json"

	HttpMethodNotAllowedError HttpErrors = "method not allowed"
)

// RespondWithData writes the result to the http.ResponseWriter with a 200 status code.
func RespondWithData(w http.ResponseWriter, result interface{}) {
	w.Header().Set(contentTypeHeader, jsonContentType)
	respondWithJSON(w, http.StatusOK, result)
}

// RespondWithStatus writes the status code to the http.ResponseWriter.
func RespondWithStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

// RespondWithError logs the error and writes it to the http.ResponseWriter with the given status code.
func RespondWithError(w http.ResponseWriter, code int, err error) {
	log.Printf("error: %s", err.Error())
	log.Println()
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}

// respondWithJSON writes the payload to the http.ResponseWriter with the given status code.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Println(err.Error())
		RespondWithStatus(w, http.StatusInternalServerError)
	}
}

// ValidateHttpMethod checks if the request method is the expected method.
func ValidateHttpMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		RespondWithStatus(w, http.StatusMethodNotAllowed)
		return false
	}
	return true
}
