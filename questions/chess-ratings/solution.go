package main

import (
	"fmt"
)

func solution(games [][]int, n int) [][]int {
	mapWinnerLosers := make(map[int][]int)
	wins := make([]int, n)
	losses := make([]int, n)
	for i := 0; i < n; i++ {
		mapWinnerLosers[i] = []int{}
	}
	for _, g := range games {
		winner, loser := g[0]-1, g[1]-1
		mapWinnerLosers[winner] = append(mapWinnerLosers[winner], loser)
	}
	for winner := range mapWinnerLosers {
		defeated := make(map[int]bool)
		toVisit := []int{winner}
		for len(toVisit) > 0 {
			w := toVisit[0]
			for _, l := range mapWinnerLosers[w] {
				if !defeated[l] {
					toVisit = append(toVisit, l)
					defeated[l] = true
					losses[l]++
				}
			}
			toVisit = toVisit[1:]
		}
		wins[winner] = len(defeated)
	}
	positions := [][]int{}
	for i := range wins {
		if wins[i]+losses[i] == n-1 {
			positions = append(positions, []int{i + 1, losses[i] + 1})
		}
	}
	return positions
}

func main() {
	fmt.Println(solution([][]int{{4, 2}, {4, 3}, {3, 2}, {1, 2}, {2, 5}, {3, 1}}, 5))
}
