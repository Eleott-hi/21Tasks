package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	val      int
	coins    []int
	expected []int
}{
	{13, []int{1, 5, 10}, []int{10, 1, 1, 1}},         // Normal case
	{23, []int{1, 5, 10}, []int{10, 10, 1, 1, 1}},     // Normal case
	{0, []int{1, 5, 10}, []int{}},                     // Zero value
	{7, []int{1, 3, 4, 7, 13, 15}, []int{7}},          // Exotic denominations
	{7, []int{13, 15}, []int{}},                       // Value smaller than smallest coin
	{23, []int{5, 10, 1}, []int{10, 10, 1, 1, 1}},     // Unsorted denominations
	{23, []int{10, 1, 5, 10}, []int{10, 10, 1, 1, 1}}, // Duplicate denominations
	{23, []int{}, []int{}},                            // Empty denominations
}

func TestMinCoins(t *testing.T) {
	for _, test := range tests {
		result := minCoins(test.val, test.coins)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test failed for value %d with coins %v, expected %v, but got %v", test.val, test.coins, test.expected, result)
		}
	}
}

func TestMinCoins2(t *testing.T) {
	for _, test := range tests {
		result := minCoins2(test.val, test.coins)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For value %d with coins %v, expected %v, but got %v", test.val, test.coins, test.expected, result)
		}
	}
}
