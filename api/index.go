package api

import (
	"encoding/json"
	"net/http"
)

func Pulls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]any{
		"message": "100",
	}
	marshal, _ := json.Marshal(response)
	_, _ = w.Write(marshal)
}
