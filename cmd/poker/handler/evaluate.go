package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"regexp"
)

var cardsRegexp = regexp.MustCompile(`^cards\[\s*]$`)

type evaluateRequest struct {
	Cards []string `validate:"required"`
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
	var (
		req = evaluateRequest{
			Cards: h.parseReq(r.URL.Query()),
		}
	)

	if err := h.validate.StructCtx(r.Context(), req); err != nil {
		writeErr(w, err)
		return
	}

	//TODO Add evaluation handling and write response.
}

func (h *EvaluateHandHandler) parseReq(uVal url.Values) []string {

	var res = make([]string, 0)
	for k := range uVal {
		if cardsRegexp.MatchString(k) {
			res = append(res, uVal.Get(k))
		}
	}
	return res
}
