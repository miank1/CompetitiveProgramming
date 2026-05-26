package main

import (
	"fmt"
)

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	lastDigit_a := a % 10
	lastDigit_b := b % 10

	fmt.Println(lastDigit_a + lastDigit_b)
}
