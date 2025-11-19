package link

import (
	"fmt"
	"net/http"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("GET /link/{hash}", handler.Get())
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) Fetch() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLink, 201)
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
		hash := r.PathValue("hash")
		res.Json(w, nil, 200);
	}
}
