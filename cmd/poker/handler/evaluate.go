package handler

import (
	"encoding/json"
	"fmt"
	"github.com/devandreyl/go-poker-hands-evaluator/internal/holdem"
	pokererr "github.com/devandreyl/go-poker-hands-evaluator/pkg/error"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type evaluateRequest struct {
	Hands holdem.Hands `json:"hands" validate:"required"`
}

type evaluateResponse struct { //TODO create response from EvaluateResult
	TestResult string
}

type EvaluateHandHandler struct {
	router   *mux.Router
	validate *validator.Validate
}

func NewEvaluateHandler(router *mux.Router, validate *validator.Validate) EvaluateHandHandler {
	return EvaluateHandHandler{
		router:   router,
		validate: validate,
	}
}

func (h *EvaluateHandHandler) Register() {
	h.router.HandleFunc("/evaluate-hand", h.evaluateHand).
		Methods(http.MethodPost)
}

func (h *EvaluateHandHandler) evaluateHand(w http.ResponseWriter, r *http.Request) {
	var req evaluateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, pokererr.Wrap(err, pokererr.CodeApiDecoderError, nil))
		return
	}

	if err := h.validate.StructCtx(r.Context(), req); err != nil {
		writeErr(w, err)
		return
	}

	result, err := holdem.EvaluateAndCompareHands(req.Hands)
	if err != nil {
		writeErr(w, err)
		return
	}
	fmt.Println(result)

	write(w, http.StatusOK, evaluateResponse{TestResult: "ok"})
}
