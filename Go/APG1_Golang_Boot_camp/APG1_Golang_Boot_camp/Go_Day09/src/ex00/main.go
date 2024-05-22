package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepSort(nums []int) <-chan int {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			ch <- n
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	nums := []int{5, 3, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1}
	ch := sleepSort(nums)

	res := make([]int, 0, len(nums))
	for num := range ch {
		res = append(res, num)
	}
	fmt.Println(res)
}
