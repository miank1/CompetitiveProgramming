// https://codeforces.com/problemset/problem/282/A

package main

import "fmt"

func main() {
	var (
		str   string
		count int
		incr  int
	)

	fmt.Scan(&incr)

	for i := 0; i < incr; i++ {
		fmt.Scan(&str)
		if str[1] == '+' {
			count++
		} else {
			count--
		}
	}

	fmt.Println(count)

}
