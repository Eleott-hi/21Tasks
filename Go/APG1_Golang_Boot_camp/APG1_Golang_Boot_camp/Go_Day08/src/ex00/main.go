package ex00

import (
	"errors"
	"unsafe"
)

func GetElement(arr []int, idx int) (int, error) {
	// Check if the slice is empty
	if len(arr) == 0 {
		return 0, errors.New("empty slice")
	}

	// Check if the index is negative
	if idx < 0 {
		return 0, errors.New("negative index")
	}

	// Check if the index is out of bounds
	if idx >= len(arr) {
		return 0, errors.New("index out of bounds")
	}

	// Use unsafe package to perform pointer arithmetic
	ptr := unsafe.Pointer(&arr[0])
	offset := uintptr(idx) * unsafe.Sizeof(arr[0])
	element := *(*int)(unsafe.Pointer(uintptr(ptr) + offset))

	return element, nil
}
