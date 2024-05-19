package main

import (
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

func grabPresents(presents []Present, capacity int) []Present {
	n := len(presents)
	if n == 0 || capacity == 0 {
		return []Present{}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if presents[i-1].Size <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-presents[i-1].Size]+presents[i-1].Value)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	result := []Present{}
	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			result = append(result, presents[i-1])
			w -= presents[i-1].Size
		}
	}

	return result
}

func main() {
	presents := []Present{
		{Value: 60, Size: 10},
		{Value: 100, Size: 20},
		{Value: 120, Size: 30},
	}
	capacity := 50

	selectedPresents := grabPresents(presents, capacity)
	fmt.Println("Selected Presents:")
	for _, p := range selectedPresents {
		fmt.Printf("Value: %d, Size: %d\n", p.Value, p.Size)
	}
}
