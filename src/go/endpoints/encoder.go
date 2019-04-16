package endpoints

import (
	"encoding/json"
	"net/http"
)

func sendResponse(rw *http.ResponseWriter, resp interface{}) {
	(*rw).WriteHeader(http.StatusOK)
	if resp == nil {
		return
	}

	(*rw).Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(*rw).Encode(resp); err != nil {
		http.Error(*rw, err.Error(), 500)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
