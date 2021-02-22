package main

import "fmt"

func main() {
	/*
	Algorithm Description
		- Imagine a series of a number from 1 to 15. (1 2 3 4 5 6 7 8 9 10 11 12 13 14 15)
		- Fizz and Buzz refer to any number that's a multiple of 3 and 5 respectively. In other words, if a number is divisible
		  by 3, it is substituted with fizz; if a number is divisible by 5, it is substituted with buzz. If a number is simultaneously
		  a multiple of 3 AND 5, the number is replaced with "fizz buzz." In essence, it emulates the famous children game "fizz buzz".
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, value := range numbers {
		if value%3 == 0 && value%5 == 0 {
			fmt.Println(value, " is FizzBuzz")
		} else if value%3 == 0 {
			fmt.Println(value, " is Fizz")
		} else if value%5 == 0 {
			fmt.Println(value, " is Buzz")
		} else {
			fmt.Println(value, " is not Fizz || Buzz || FizzBuzz")
		}
	}
}
