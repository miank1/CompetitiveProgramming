// https://codeforces.com/problemset/problem/4/A
package main

import "fmt"

func main() {
	var w int
	_, err := fmt.Scan(&w)
	if err != nil {
		fmt.Println("error reading the input")
	}

	if w == 2 {
		fmt.Println("NO")
		return
	}
	if w%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
