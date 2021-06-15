package calculator

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumber() (int, int) {
	/*
		@Documentation rand:

		Top-level functions, such as Float64 and Int, use a default shared Source
		that produces a deterministic sequence of values each time a program is run.
		Use the Seed function to initialize the default Source if different behavior
		is required for each run.
	*/

	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 999

	n1 := rand.Intn(max-min) + min
	n2 := rand.Intn(max-min) + min

	return n1, n2
}

func printCalc() {
	n1, n2 := generateNumber()
	fmt.Printf("[PROCCESS 1 - CALCULATOR] Result: %v + %v =  %v \n", n1, n2, n1+n2)
	fmt.Printf("---------------------------------------- \n")
}

func LoopCalcs() {
	for x := 0; x < 15; x++ {
		printCalc()
	}
}
