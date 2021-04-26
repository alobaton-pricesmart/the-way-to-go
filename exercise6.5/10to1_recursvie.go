package main

import "fmt"

func main() {
	printNto1Recursive(10)
}

func printNto1Recursive(n int) {
	if n < 1 {
		return
	}
	defer fmt.Println(n)
	printNto1Recursive(n - 1)
}
