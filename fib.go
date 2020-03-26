package main

import (
	"strconv"
	"fmt"
)

// create a function that returns the nth term of the Fibonacci series
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	nth := Fib(n-1) + Fib(n-2)
	return nth
}

func main() {
	n := Fib(7)
	fmt.Println(strconv.Itoa(n))
}
