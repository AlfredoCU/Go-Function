package main

import "fmt"

// SwapNumber Function
func swapNumber(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	fmt.Print("a: ", *a," b: ", *b)
}

// Main Function
func main() {
	a := 5
	b := 7
	fmt.Println("a:", a,"b:", b)
	swapNumber(&a, &b)
}