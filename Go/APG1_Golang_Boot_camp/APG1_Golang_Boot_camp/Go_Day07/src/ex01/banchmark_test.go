package main

import (
	"math/rand"
	"testing"
	"time"
)

// Generate a slice of random coin denominations
func generateRandomCoins(n int) []int {
	rand.Seed(time.Now().UnixNano())
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		coins[i] = rand.Intn(100000) + 1 // Ensure non-zero denominations
	}
	return coins
}

func BenchmarkMinCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins(100000, []int{1, 5, 10, 25, 50, 100})
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins2(100000, []int{1, 5, 10, 25, 50, 100})
	}
}

func BenchmarkMinCoins2Unsorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins2(100000, []int{100, 50, 25, 10, 5, 1})
	}
}

func BenchmarkMinCoins2Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins2Optimized(100000, []int{1, 5, 10, 25, 50, 100})
	}
}

var largeCoins = generateRandomCoins(10000) // Adjust size as needed

// New benchmark functions with a large set of random coin denominations
func BenchmarkMinCoinsLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins(100000, largeCoins)
	}
}

func BenchmarkMinCoins2Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins2(100000, largeCoins)
	}
}

func BenchmarkMinCoins2UnsortedLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins2(100000, largeCoins)
	}
}

func BenchmarkMinCoins2OptimizedLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCoins2Optimized(100000, largeCoins)
	}
}
