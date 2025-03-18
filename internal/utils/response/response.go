package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status int
	Error string
}

func WriteJson(w http.ResponseWriter, status int, payload any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}

func GeneralError(err error) ErrorResponse {
	return ErrorResponse {
		Status: 400,
		Error: err.Error(),
	}
}

func ValidationError(errors validator.ValidationErrors) ErrorResponse {

	var errorsStr []string

	for _, err := range errors {
		errorsStr = append(errorsStr, fmt.Errorf("Field %s is required", err.Field()).Error())
	}

	return ErrorResponse {
		Status: 400,
		Error: strings.Join(errorsStr, ","),
	}
}