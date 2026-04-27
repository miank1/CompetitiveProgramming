package main

import "fmt"

func main() {

	var i, j, k int

	_, err := fmt.Scan(&i, &j, &k)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	max := i
	min := i

	if j < min {
		min = j
	}

	if k < min {
		min = k
	}

	if j > max {
		max = j
	}
	if k > max {
		max = k
	}

	fmt.Printf("%d %d", min, max)

}
