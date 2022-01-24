package main

import (
	"fmt"
	"time"

	"bruteForce/solution"
)

func main() {
	const (
		start = 4
		end   = 13
	)
	for i := start; i <= end; i++ {
		s := time.Now()
		ans := solution.SolveQueen(i)
		e := time.Now()
		fmt.Printf("n = %2d, answer = %8d, time elapsed = %.3f ms\n", i, ans, float64(e.Sub(s).Microseconds())/1000.0)
	}

	fmt.Println("------------------------------------------------------------------")

	for i := start; i <= end; i++ {
		s := time.Now()
		ans := solution.SolveQueenMulti(i)
		e := time.Now()
		fmt.Printf("n = %2d, answer = %8d, time elapsed = %.3f ms\n", i, ans, float64(e.Sub(s).Microseconds())/1000.0)
	}
}
