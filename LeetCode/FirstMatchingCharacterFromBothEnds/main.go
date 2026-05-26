package main

func firstMatchingIndex(s string) int {

	j := len(s) - 1
	r := []rune(s)

	result := -1

	for i := 0; i <= j; i++ {
		if r[i] == r[j] {
			result = i
			break
		}
		j--
	}

	return result
}

func main() {
	s := "abcdab"
	firstMatchingIndex(s)
}
