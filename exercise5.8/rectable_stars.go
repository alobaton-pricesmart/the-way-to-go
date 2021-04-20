package main

import (
	"fmt"
	"strings"
)

func main() {
	rectangleStars(10, 20)
}

func rectangleStars(w, h int) {
	for i := 0; i < h; i++ {
		var b strings.Builder
		for j := 0; j < w; j++ {
			b.WriteString("*")
		}
		fmt.Println(b.String())
	}
}
