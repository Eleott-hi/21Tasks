package main

import (
    "container/heap"
    "fmt"
)

// Present struct
type Present struct {
    Value int
    Size  int
}

// PresentHeap implements heap.Interface and holds Presents
type PresentHeap []Present

func (h PresentHeap) Len() int {
    return len(h)
}

func (h PresentHeap) Less(i, j int) bool {
    if h[i].Value == h[j].Value {
        return h[i].Size < h[j].Size
    }
    return h[i].Value > h[j].Value
}

func (h PresentHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *PresentHeap) Push(x interface{}) {
    *h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[0 : n-1]
    return item
}

func getNCoolestPresents(presents []Present, n int) ([]Present, error) {
    if n < 0 || n > len(presents) {
        return nil, fmt.Errorf("invalid value for n: %d", n)
    }

    h := &PresentHeap{}
    heap.Init(h)

    for _, present := range presents {
        heap.Push(h, present)
    }

    coolest := make([]Present, n)
    for i := 0; i < n; i++ {
        coolest[i] = heap.Pop(h).(Present)
    }

    return coolest, nil
}

func main() {
    presents := []Present{
        {Value: 5, Size: 1},
        {Value: 4, Size: 5},
        {Value: 3, Size: 1},
        {Value: 5, Size: 2},
    }

    n := 2
    coolestPresents, err := getNCoolestPresents(presents, n)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Coolest Presents:", coolestPresents)
    }
}
