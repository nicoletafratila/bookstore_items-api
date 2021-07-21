package http_utils

import (
	"encoding/json"
	"fmt"
	"github.com/nicoletafratila/bookstore_utils-go/rest_errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if result := json.NewEncoder(w).Encode(body); result != nil {
		fmt.Println("Error json: " + result.Error())
	}
}

func RespondError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJson(w, err.Status(), err)
}
