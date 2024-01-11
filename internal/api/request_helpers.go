package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, err string) {
	RespondWithJSON(w, code, map[string]string{"error": fmt.Sprintf("%+v", err)})
}

func GetParamsFromRequestBody[T interface{}](structBody T, r *http.Request) (T, error) {
	decoder := json.NewDecoder(r.Body)

	params := structBody
	err := decoder.Decode(&params)
	if err != nil {
		return structBody, errors.New("couldn't decode parameters")
	}

	return params, nil
}

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"status": "ok"}
	RespondWithJSON(w, 200, res)
}

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 500, "Internal Server Error")
}
