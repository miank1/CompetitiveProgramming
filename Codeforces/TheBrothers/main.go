package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// First line
	scanner.Scan()
	first := scanner.Text()

	// Second line
	scanner.Scan()
	second := scanner.Text()

	part1 := strings.Split(first, " ")

	lastName_1 := part1[len(part1)-1]

	part2 := strings.Split(second, " ")

	lastName_2 := part2[len(part2)-1]

	if lastName_1 == lastName_2 {
		fmt.Println("ARE Brothers")
	} else {
		fmt.Println("NOT")
	}
}
