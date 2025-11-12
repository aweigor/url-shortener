package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/pkg/res"

	"github.com/go-playground/validator/v10"
)

const (
  ErrEmailNotDefined  = "email is required"
	ErrPasswordNotDefined  = "password is required"
	ErrEmailNotValid = "email is invalid"
	ErrJsonNotValid = "json body parser error"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		// read body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, ErrJsonNotValid, 400)
			return
		}
		validate := validator.New()
		err = validate.Struct(payload )
		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200);
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		fmt.Println("OK")
	}
}