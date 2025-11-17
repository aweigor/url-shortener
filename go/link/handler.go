package link

import (
	"fmt"
	"net/http"
	"url-shortener/pkg/res"
)



type LinkHandler struct {}

type LinkHandlerDeps struct {}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{}
	router.HandleFunc("GET /link/{hash}", handler.Get())
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) Get() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		res.Json(w, nil, 200);
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		res.Json(w, nil, 200);
	}
}
	
func (handler *LinkHandler) Update() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
		res.Json(w, nil, 200);
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		
		res.Json(w, nil, 200);
	}
}
