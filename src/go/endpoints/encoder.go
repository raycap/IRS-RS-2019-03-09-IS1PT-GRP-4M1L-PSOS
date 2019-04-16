package endpoints

import (
	"encoding/json"
	"net/http"
)

func sendResponse(rw http.ResponseWriter, resp interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if resp == nil {
		return
	}
	if err := json.NewEncoder(rw).Encode(resp); err != nil {
		http.Error(rw, err.Error(), 500)
	}
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
