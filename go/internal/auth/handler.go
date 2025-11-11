package auth

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/pkg/res"
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
		fmt.Println("OK")
		res.Json(w, LoginResponse{ Token: "123" }, 200);
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		fmt.Println("OK")
	}
}