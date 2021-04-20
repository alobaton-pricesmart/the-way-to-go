package main

import (
	"fmt"
	"strings"
)

func main() {
	forChar(25, "G")
}

func forChar(n int, c string) {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(c)
		fmt.Println(b.String())
	}
}
