package main

import "fmt"

func main() {
	forLoop(15)
	gotoLoop(15)
}

func forLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func gotoLoop(n int) {
	i := 0
start:
	fmt.Println(i)
	i++
	if i < n {
		goto start
	}
}
