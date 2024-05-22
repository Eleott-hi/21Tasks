package main

import (
	// "fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Call the function with test data
	testData := []struct {
		val   int
		coins []int
	}{
		{100000000, []int{1, 5, 10, 25, 50, 100}},
		{100000000, []int{1, 3, 4, 7, 13, 15}},
	}

	for _, td := range testData {
		minCoins(td.val, td.coins)
	}

	for _, td := range testData {
		minCoins2(td.val, td.coins)
	}

	for _, td := range testData {
		minCoins2Optimized(td.val, td.coins)
	}
}
