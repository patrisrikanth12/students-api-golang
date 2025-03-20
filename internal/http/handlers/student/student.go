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
	"github.com/patrisrikanth12/students-api-golang/internal/storage"
)

func Create(storage storage.Storage) http.HandlerFunc {
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

		id, err := storage.CreateStudent(student.Name, student.Email, student.Mobile)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Created student with", slog.Int64("id", id))

		response.WriteJson(w, http.StatusCreated, map[string] string {"Success": "OK"})
	}
}