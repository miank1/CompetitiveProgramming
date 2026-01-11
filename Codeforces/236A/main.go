package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)


	m1 := make(map[string]int)

	for _, v := range n {
		m1[string(v)]++
	}

	if len(m1) % 2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}