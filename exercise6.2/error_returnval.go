package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	s, e := mySqrtUnnamed(9)
	if e != nil {
		panic(e)
	}
	fmt.Println(s)

	s, e = mySqrtNamed(9)
	if e != nil {
		panic(e)
	}
	fmt.Println(s)

}

func mySqrtUnnamed(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("a must be positive")
	}

	return math.Sqrt(a), nil
}

func mySqrtNamed(a float64) (s float64, e error) {
	if a < 0 {
		e = errors.New("a must be positive")
		return
	}

	s = math.Sqrt(a)
	return
}
