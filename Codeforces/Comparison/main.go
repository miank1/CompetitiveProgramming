package main

import "fmt"

func main() {
	var m, n int

	_, err := fmt.Scan(&m, &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if m > n {
		fmt.Println("Right")
	} else if m < n {
		fmt.Println("Wrong")
	}
}
