package main

import (
	"fmt"
	"time"

	"./ga"
)

func main() {
	now := time.Now().UnixNano()
	modeller := &ga.ChromosomeModellerNoOp{}
	gaSolver := ga.New(10, 2, 10, 0.05, modeller)
	fmt.Println(gaSolver.Solve())
	fmt.Println(time.Now().UnixNano() - now)
}
