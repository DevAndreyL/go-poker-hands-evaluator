package handler

import (
	"encoding/json"
	pokererr "github.com/devandreyl/go-poker-hands-evaluator/pkg/error"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type errorResponse struct {
	Error error `json:"error"`
}

func write(w http.ResponseWriter, status int, resp any) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Panic(err)
	}
}

func writeErr(w http.ResponseWriter, err error) {

	if conv, ok := err.(validator.ValidationErrors); ok {
		data := make(pokererr.Data)
		for _, v := range conv {
			data[v.StructField()] = v.Error()
		}

		write(w, http.StatusBadRequest, errorResponse{
			Error: pokererr.NewError(pokererr.CodeValidationError, data),
		})

		return
	}

	if conv, ok := err.(*pokererr.Error); ok {
		write(w, http.StatusInternalServerError, errorResponse{
			Error: conv,
		})

		return
	}

	write(w, http.StatusInternalServerError, errorResponse{
		Error: pokererr.Wrap(err, pokererr.CodeUnknown, nil),
	})
}
