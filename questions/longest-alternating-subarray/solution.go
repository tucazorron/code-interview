package main

import (
	"fmt"
)

func solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	maxLen, currentLen := 1, 1

	for i := 1; i < len(A); i++ {
		if i >= 2 && ((A[i] > A[i-1] && A[i-1] < A[i-2]) || (A[i] < A[i-1] && A[i-1] > A[i-2])) {
			currentLen++
		} else if A[i] != A[i-1] {
			currentLen = 2
		} else {
			currentLen = 1
		}

		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

func main() {
	fmt.Println(solution([]int{1, 2, 1, 2}))          // Output: 4
	fmt.Println(solution([]int{1, 2, 3, 2, 1, 2, 1})) // Output: 4
	fmt.Println(solution([]int{1, 1, 1, 1, 1}))       // Output: 1
	fmt.Println(solution([]int{2, 3, 2, 3, 2, 3, 2})) // Output: 7
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7})) // Output: 2
	fmt.Println(solution([]int{1, 2, 1, 3, 1, 3, 1})) // Output: 5
}
