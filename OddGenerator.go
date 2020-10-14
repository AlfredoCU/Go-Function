package main

import "fmt"

// OddGenerator return anonymous function and anonymous function return integer
func oddGenerator() func() int {
	i := 1 // Initialize only once.
	return func() int {
		var odd = i
		i += 2
		return odd
	}
}

// Main Function
func main() {
	newOdd := oddGenerator()
	for i := 1; i <= 50; i++ {
		fmt.Print(newOdd(), " ")
	}
}