package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5))
}

func fibonacci(n int) (l, r int) {
	if n <= 1 {
		l = 0
		r = 1
		return
	}

	a, b := fibonacci(n - 1)
	c, d := fibonacci(n - 2)

	return a + b, c + d
}
