package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

func main() {
	num1, num2 := 5.6, 9.5
	fmt.Println(add(num1, num2))
	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1, w2))
}

func multiple(a, b string) (string, string) {
	return a, b
}
