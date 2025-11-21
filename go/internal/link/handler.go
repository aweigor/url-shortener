package link

import (
	"net/http"
	"strconv"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"

	"gorm.io/gorm"
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
	router.HandleFunc("GET /link/{hash}", handler.Fetch())
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
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
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
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ ID: uint(id) },
			Url: body.Url,
			Hash: body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, link, 201)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
		}

		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, nil, 200)
	}
}
