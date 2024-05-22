package ex02_test

import (
	"example.com/ex02"
	"fmt"
)

// ExampleMinCoins demonstrates the usage of the minCoins function to find
// the minimum number of coins needed to make a given value using a set of
// available coins.
func ExampleMinCoins() {
	val := 17
	coins := []int{1, 2, 5, 10, 20}
	result := ex02.MinCoins(val, coins)
	fmt.Println(result)
	// Output: [10 5 2]
}

// ExampleMinCoins2 demonstrates the usage of the minCoins2 function to find
// the minimum number of coins needed to make a given value using a set of
// available coins.
func ExampleMinCoins2() {
	val := 17
	coins := []int{1, 2, 5, 10, 20}
	result := ex02.MinCoins2(val, coins)
	fmt.Println(result)
	// Output: [10 5 2]
}

// ExampleMinCoins2Optimized demonstrates the usage of the minCoins2Optimized
// function to find the minimum number of coins needed to make a given value
// using a set of available coins.
func ExampleMinCoins2Optimized() {
	val := 17
	coins := []int{1, 2, 5, 10, 20}
	result := ex02.MinCoins2Optimized(val, coins)
	fmt.Println(result)
	// Output: [10 5 2]
}
