package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/dto/response"
	"github.com/renaldyhidayatt/redditgoent/services"
)

type subRedditHandler struct {
	service services.SubRedditService
}

func NewSubRedditHandler(service services.SubRedditService) *subRedditHandler {
	return &subRedditHandler{service: service}
}

func (h *subRedditHandler) Create(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}
	var subredditRequest request.SubRedditRequest

	validate := validator.New()
	err := validate.Struct(subredditRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	res, err := h.service.Create(email, subredditRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *subRedditHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *subRedditHandler) FindById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	res, err := h.service.FindById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}
