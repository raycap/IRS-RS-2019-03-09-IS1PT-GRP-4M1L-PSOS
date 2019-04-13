package main

import (
	"fmt"
	"net/http"

	"./endpoints"
)

type GaParams struct {
	PopSize     int64 `json:"popSize"`
	EliteSize   int64 `json:"eliteSize"`
	Generations int64 `json:"generations"`
	QuickScan   bool  `json:"quickScan"`
	UseCon      bool  `json:"useCon"`
}

func main() {
	http.HandleFunc("/testGaParam", endpoints.TestGaParam)

	http.HandleFunc("/solve", endpoints.Solve)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("err running server %s", err)
	}
}
