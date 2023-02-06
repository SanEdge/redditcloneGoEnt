package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/dto/response"
	"github.com/renaldyhidayatt/redditgoent/security"
	"github.com/renaldyhidayatt/redditgoent/services"
)

type authHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest request.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(registerRequest)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.service.Register(registerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var authRequest request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(authRequest)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.service.Login(authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if res.ID > 0 {
		hashPwd := res.Password
		pwd := authRequest.Password

		hash := security.VerifyPassword(hashPwd, pwd)

		fmt.Println(res.Email)

		if hash == nil {
			token, err := security.GenerateToken(res.Email)

			if err != nil {
				response.ResponseError(w, http.StatusInternalServerError, err)
				return
			}

			response.ResponseToken(w, "Login Berhasil", token, res, http.StatusOK)
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("password tidak sesuai"))
			return
		}
	}
}

func (h *authHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	res, err := h.service.VerifyToken(token)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan verify token", res, http.StatusOK)
}
