// Synopsis:
// This program provides functions to calculate the minimum number of coins
// required to make a given value using a set of available coins. It includes
// different implementations of the algorithm, each with its optimizations.
//
// Overview:
// The main goal of this program is to solve the minimum coin change problem,
// which involves finding the minimum number of coins needed to make a given
// value using a set of available coin denominations. The program provides
// alternative implementations of the algorithm, showcasing different approaches
// and optimizations.
//
// Instructions to generate documentation:
//
// 1. Run the following command:
//
//    go doc -all > documentation.txt
//
// This command generates documentation comments for all the package members
// and redirects the output to a text file named `documentation.txt`.
//
// 5. Optionally, you can use `godoc` to generate HTML documentation:
//
//    godoc -http=:6060
//
// This command starts a local web server and serves documentation for your Go packages.
//
// 6. Open your web browser and navigate to `http://localhost:6060/pkg/` to view the documentation.
//
package ex02

import (
	"sort"
)


// MinCoins returns the minimum number of coins needed to make the given value
// using the provided coins. This function uses a greedy algorithm to achieve
// the result.
func MinCoins(val int, coins []int) []int {
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

// MinCoins2 is an alternate implementation of minCoins using a different approach.
// It sorts the coins in descending order and iterates through them to find the
// minimum number of coins needed.
func MinCoins2(val int, coins []int) []int {
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

// MinCoins2Optimized is an optimized version of minCoins2.
// It preallocates the result slice based on the maximum number of coins needed,
// which is val/coins[0], where coins[0] is the largest coin value.
func MinCoins2Optimized(val int, coins []int) []int {
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

