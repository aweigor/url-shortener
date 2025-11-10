package main

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/heartbeat"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	heartbeat.NewHeartbeatHandler(router)
	
	server := http.Server{ Addr: ":8081", Handler: router }

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
