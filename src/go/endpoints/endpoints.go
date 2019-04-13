package endpoints

import (
	"net/http"

	"fmt"
	"time"

	"../services"
	"../services/dto"
)

func TestGaParam(rw http.ResponseWriter, req *http.Request) {
	params, err := parsePOSTRequest(rw, req, &dto.GaParams{})
	if err != nil {
		return
	}
	reqParams := params.(*dto.GaParams)

	now := time.Now().UnixNano()
	resp, err := services.ResultWithGaParam(reqParams)
	fmt.Printf("time taken to solve the request : %d ms\n", (time.Now().UnixNano()-now)/1000000)

	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	sendResponse(rw, resp)
}

func Solve(rw http.ResponseWriter, req *http.Request) {
	params, err := parsePOSTRequest(rw, req, &dto.RequestPayload{})
	if err != nil {
		return
	}
	reqParams := params.(*dto.RequestPayload)

	now := time.Now().UnixNano()
	resp, err := services.SolveWithFixtures(reqParams)
	fmt.Printf("time taken to solve the request : %d ms\n", (time.Now().UnixNano()-now)/1000000)

	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	sendResponse(rw, resp)
}
