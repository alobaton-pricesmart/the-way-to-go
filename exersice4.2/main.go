package main

import "fmt"

var a = "G"

func main() {
	n() // Should print G
	m() // Should print O
	n() // Should print O
}
func n() {
	fmt.Println(a)
}
func m() {
	a = "O"
	fmt.Println(a)
}
