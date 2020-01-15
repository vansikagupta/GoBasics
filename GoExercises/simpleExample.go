package main

import (
	"fmt"
	"math"
	"math/rand"
)

func math_func(n float64) {
	fmt.Println("square root of %v is :", n, math.Sqrt(n))
	fmt.Println("random number :", rand.Intn(40))
}

func string_func(a, b string) (string, string) {
	return b, a
}
func main() {
	fmt.Println("Welcome to Go!!")
	math_func(15)
	fmt.Println(string_func("hello", "there"))
}
