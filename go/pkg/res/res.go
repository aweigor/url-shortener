package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}