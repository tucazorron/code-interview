package main

import (
	"fmt"
)

func solution(nums []int) int {
	missing := 0
	for i, num := range nums {
		missing += i + 1 - num
	}
	return missing + len(nums) + 1
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7, 9})) // 8
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})) // 10
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 10})) // 9
	fmt.Println(solution([]int{2, 3})) // 1
}
