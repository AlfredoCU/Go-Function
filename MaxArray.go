package main

import "fmt"

func MaxArray(args ...int) int {
	n := 0
	for i,_ := range args {
		if  n < args[i] {
			n = args[i]
		}
	}
	return n
}

func main() {
	s := []int{-1, 2, 5, 1, 7, 3, 6, -10}
	fmt.Print(MaxArray(s...))
}