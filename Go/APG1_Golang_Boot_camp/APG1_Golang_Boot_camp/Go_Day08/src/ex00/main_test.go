package ex00_test

import (
	"testing"

	"example.com/ex00"
)

var arr = []int{10, 20, 30, 60, 50}

func TestGetElementValid(t *testing.T) {
	index := 3
	element, err := ex00.GetElement(arr, index)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	expected := 60
	if element != expected {
		t.Errorf("Expected element at index %d to be %d, got %d", index, expected, element)
	}
}

func TestGetElementEmptySlice(t *testing.T) {
	emptySlice := []int{}
	_, err := ex00.GetElement(emptySlice, 0)
	if err == nil {
		t.Error("Expected error for empty slice, got nil")
	}
}

func TestGetElementNegativeIndex(t *testing.T) {
	_, err := ex00.GetElement(arr, -1)
	if err == nil {
		t.Error("Expected error for negative index, got nil")
	}
}

func TestGetElementIndexOutOfBounds(t *testing.T) {
	_, err := ex00.GetElement(arr, 10)
	if err == nil {
		t.Error("Expected error for index out of bounds, got nil")
	}
}
