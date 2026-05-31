package main

import "fmt"

func main() {

	var a, b, c, d int64

	fmt.Scan(&a, &b, &c, &d)

	if a+b+c == d ||
		a+b-c == d ||
		a+b*c == d ||
		a-b+c == d ||
		a-b-c == d ||
		a-b*c == d ||
		a*b+c == d ||
		a*b-c == d ||
		a*b*c == d ||
		(a+b)*c == d ||
		(a-b)*c == d ||
		(a*b)*c == d {

		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
