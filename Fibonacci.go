package main

import "fmt"

// Fibonacci Function
func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

// Main Function
func main() {
	fmt.Print(fibonacci(10))
}
