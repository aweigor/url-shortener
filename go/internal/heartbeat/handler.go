package heartbeat

import (
	"fmt"
	"net/http"
)

type HeartbeatHandler struct {}

func NewHeartbeatHandler(router *http.ServeMux) {
	handler := &HeartbeatHandler{}
	router.HandleFunc("/heartbeat", handler.Hello())
}

func (handler *HeartbeatHandler) Hello() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		fmt.Println("OK")
	}
}