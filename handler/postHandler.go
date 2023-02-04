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
	"github.com/renaldyhidayatt/redditgoent/interfaces"
	"github.com/renaldyhidayatt/redditgoent/services"
)

type postHandler struct {
	service interfaces.IPostService
}

func NewPostHandler(service services.PostService) *postHandler {
	return &postHandler{service: service}
}

func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}
	var postRequest request.PostRequest

	validate := validator.New()
	err := validate.Struct(postRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.service.Create(email, postRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}

}

func (h *postHandler) FindAll(w http.ResponseWriter, r *http.Request) {

	res, err := h.service.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *postHandler) FindById(w http.ResponseWriter, r *http.Request) {
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

func (h *postHandler) FindPostBySubreddit(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.service.FindPostBySubreddit(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)

}

func (h *postHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}

	res, err := h.service.FindByEmail(email)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}
