package search

import "testing"

func TestBinarySearchOddLengthSlice(t *testing.T) {
	list := []int{1, 3, 5, 7, 9}
	if BinarySearch(list, 3) != 1 {
		t.Error("BinarySearch failed")
	}
	if BinarySearch(list, -1) != -1 {
		t.Error("BinarySearch failed")
	}
	if BinarySearch(list, 9) != 4 {
		t.Error("BinarySearch failed")
	}
	if BinarySearch(list, 5) != 2 {
		t.Error("BinarySearch failed")
	}
}

func TestBinarySearchEvenLengthSlice(t *testing.T) {
	list := []int{1, 3, 5, 7, 9, 11}
	if BinarySearch(list, 3) != 1 {
		t.Error("BinarySearch failed")
	}
	if BinarySearch(list, -1) != -1 {
		t.Error("BinarySearch failed")
	}
	if BinarySearch(list, 9) != 4 {
		t.Error("BinarySearch failed")
	}
	if BinarySearch(list, 5) != 2 {
		t.Error("BinarySearch failed")
	}
}

func TestBinarySearchEmpty(t *testing.T) {
	var list []int

	if BinarySearch(list, 1) != -1 {
		t.Error("BinarySearch failed")
	}
}

func TestBinarySearchOneElementSlice(t *testing.T) {
	list := []int{1}

	if BinarySearch(list, 1) != 0 {
		t.Error("BinarySearch failed")
	}
	if BinarySearch(list, -1) != -1 {
		t.Error("BinarySearch failed")
	}
}
