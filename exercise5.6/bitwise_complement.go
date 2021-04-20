package main

import (
	"fmt"
	"strings"
)

func main() {
	bitwiseComplement(10)
}

func bitwiseComplement(n int) {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(fmt.Sprintf("%b", int64(i)))
	}
	fmt.Println(b.String())
}
