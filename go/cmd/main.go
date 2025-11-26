package main

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/heartbeat"
	"url-shortener/internal/link"
	"url-shortener/pkg/db"
	"url-shortener/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	
	router := http.NewServeMux()
	heartbeat.NewHeartbeatHandler(router)
	
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	linkRepository := link.NewLinkRepository(database)
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	mwStack := middleware.Chain(middleware.CORS, middleware.Logging)
	
	server := http.Server{ Addr: ":8081", Handler: mwStack(router) }

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}