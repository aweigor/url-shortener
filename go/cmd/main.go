package main

import (
	"fmt"
	"go/url-shortener/internal/heartbeat"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	heartbeat.NewHeartbeatHandler(router)
	
	server := http.Server{ Addr: ":8081", Handler: router }

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
