package main

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/heartbeat"
	"url-shortener/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	test := db.NewDb(conf)
	fmt.Println(test.DB.Statement.Vars...)
	router := http.NewServeMux()
	heartbeat.NewHeartbeatHandler(router)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	
	server := http.Server{ Addr: ":8081", Handler: router }

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
