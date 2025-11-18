package main

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/heartbeat"
	"url-shortener/link"
	"url-shortener/pkg/db"
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
	
	server := http.Server{ Addr: ":8081", Handler: router }

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
