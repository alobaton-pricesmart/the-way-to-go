package main

import "fmt"

func main() {
	s, m, d := multReturnUnnamed(6, 2)
	fmt.Println(s, m, d)
	s, m, d = multReturnNamed(6, 2)
	fmt.Println(s, m, d)
}

func multReturnUnnamed(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func multReturnNamed(a, b int) (s, m, d int) {
	s = a + b
	m = a * b
	d = a - b
	return
}
