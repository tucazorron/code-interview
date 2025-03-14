package main

import "fmt"

type input struct {
	array  []int
	target int
}

func solution(nums []int, target int) bool {
	diffSet := make(map[int]bool)
	for _, num := range nums {
		if diffSet[num] {
			return true
		}
		diffSet[target-num] = true
	}
	return false
}

func main() {
	tests := []input{
		{[]int{1, 1}, 2},       // true
		{[]int{0, 9, 0}, 0},    // false
		{[]int{1}, 1},          // false
		{[]int{1, 2, 4, 3}, 4}, // false

	}
	for _, test := range tests {
		result := solution(test.array, test.target)
		fmt.Printf("test: %v => %v\n", test, result)
	}
}
