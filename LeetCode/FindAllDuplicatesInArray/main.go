// 442 - https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

package main

func findDuplicates(nums []int) []int {

	m := make(map[int]int)

	duplicates := []int{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for x, v := range m {
		if v > 1 {
			duplicates = append(duplicates, x)
		}
	}

	return duplicates

}

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDuplicates(arr)

}
