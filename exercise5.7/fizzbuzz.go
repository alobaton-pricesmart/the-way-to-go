package main

import "fmt"

func main() {
	fizzbuzz(15)
}

func fizzbuzz(n int) {
	for i := 0; i < n; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println("FizzBuzz")
		}
	}
}
