package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func parsePOSTRequest(rw http.ResponseWriter, req *http.Request, defaultValue interface{}) (interface{}, error) {
	fmt.Printf("received request : method :%s, url : %s, body: %v\n", req.Method, req.URL, req.Body)
	if req.Method != "POST" {
		http.Error(rw, "request must be POST", 400)
		return defaultValue, fmt.Errorf("method error")
	}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return defaultValue, err
	}

	// Unmarshal
	copied := reflect.New(reflect.ValueOf(defaultValue).Elem().Type()).Interface()
	err = json.Unmarshal(b, copied)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return defaultValue, err
	}
	return copied, nil
}
