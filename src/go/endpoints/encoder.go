package endpoints

import (
	"encoding/json"
	"net/http"
)

func sendResponse(rw http.ResponseWriter, resp interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(resp); err != nil {
		http.Error(rw, err.Error(), 500)
	}
}
