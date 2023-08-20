package handler

import (
	"encoding/json"
	"errors"
	pokererr "github.com/devandreyl/go-poker-hands-evaluator/pkg/error"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type errorResponse struct {
	Error error `json:"error"`
}

func writeJson(w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Panic(err)
	}
}

func writeJsonErr(w http.ResponseWriter, err error) {

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		data := make(pokererr.Data)
		for _, v := range validationErrors {
			data[v.StructField()] = v.Error()
		}

		writeJson(w, http.StatusBadRequest, errorResponse{
			Error: pokererr.NewError(pokererr.CodeValidationError, data),
		})

		return
	}

	var pokerError *pokererr.Error
	if errors.As(err, &pokerError) {
		writeJson(w, http.StatusInternalServerError, errorResponse{
			Error: pokerError,
		})

		return
	}

	writeJson(w, http.StatusInternalServerError, errorResponse{
		Error: pokererr.Wrap(err, pokererr.CodeUnknown, nil),
	})
}
