package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyLetters(s string) string {

	arr := []rune(s)
	i, j := 0, len(arr)-1

	for i < j {

		if !unicode.IsLetter(arr[i]) {
			i++
			continue
		}

		if !unicode.IsLetter(arr[j]) {
			j--
			continue
		}

		// swap
		arr[i], arr[j] = arr[j], arr[i]

		i++
		j--
	}

	return string(arr)

}

func main() {
	str := "ab-cd"

	fmt.Println(reverseOnlyLetters(str))
}
