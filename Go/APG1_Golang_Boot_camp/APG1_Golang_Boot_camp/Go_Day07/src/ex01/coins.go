package main

import (
	"sort"
)

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		if val == 0 {
			break
		}
		i -= 1
	}
	return res
}

func minCoins2(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	res := make([]int, 0)
	for _, coin := range coins {
		for val >= coin {
			val -= coin
			res = append(res, coin)
		}
	}
	return res
}

func minCoins2Optimized(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	res := make([]int, val/coins[0])
	for _, coin := range coins {
		for val >= coin {
			val -= coin
			res = append(res, coin)
		}
	}
	return res
}
