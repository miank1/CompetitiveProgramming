// https://codeforces.com/problemset/problem/158/A

package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	scores := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	threshold := scores[k-1]
	count := 0

	for i := 0; i < n; i++ {
		if scores[i] >= threshold && scores[i] > 0 {
			count++
		}
	}

	fmt.Println(count)
}
