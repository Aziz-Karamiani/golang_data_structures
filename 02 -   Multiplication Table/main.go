package main

import "fmt"

func main() {
	// Multiplication Table
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%v \t", i*j)
		}
		fmt.Println()
	}
}
