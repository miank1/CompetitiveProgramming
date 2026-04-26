package main

import "fmt"

func main() {
	var m, n int

	_, err := fmt.Scan(&m, &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if m%n == 0 || n%m == 0 {
		fmt.Println("Multiples")
	} else {
		fmt.Println("No Multiples")
	}
}
