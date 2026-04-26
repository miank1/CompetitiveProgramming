// https://codeforces.com/problemset/problem/271/A

package main

import "fmt"

func allDigitsDistinct(n int) bool {
	seen := [10]bool{}

	for n > 0 {
		digit := n % 10
		if seen[digit] {
			return false
		}
		seen[digit] = true
		n /= 10
	}
	return true
}

func main() {
	var num int
	fmt.Scan(&num)

	for i := num + 1; ; i++ {
		if allDigitsDistinct(i) {
			fmt.Println(i)
			return
		}
	}
}
