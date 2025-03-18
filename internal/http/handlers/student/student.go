package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/patrisrikanth12/students-api-golang/internal/types"
	"github.com/patrisrikanth12/students-api-golang/internal/utils/response"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		
		err := json.NewDecoder(r.Body).Decode(&student) 

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validationErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validationErrs))
			return
		}

		slog.Info("Creating student") 

		response.WriteJson(w, http.StatusCreated, student)
	}
}