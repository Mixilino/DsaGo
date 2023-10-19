package search

import (
	"DsaPrimeagen/utils/slice"
	"testing"
)

func TestCrystalBallSearchJump4(t *testing.T) {
	// elem = index * 2 + 1
	list := slice.GenerateSlice(16, 1, 2)
	// check indexes 1, 0, 4, 15, 11

	if CrystalBallSearch(list, 3) != 1 {
		t.Error("BinarySearch failed. Expected 1, got ", CrystalBallSearch(list, 3))
	}
	if CrystalBallSearch(list, 1) != 0 {
		t.Error("BinarySearch failed. Expected 0, got ", CrystalBallSearch(list, 1))
	}
	if CrystalBallSearch(list, 9) != 4 {
		t.Error("BinarySearch failed. Expected 4, got ", CrystalBallSearch(list, 9))
	}
	if CrystalBallSearch(list, 31) != 15 {
		t.Error("BinarySearch failed. Expected 16, got ", CrystalBallSearch(list, 31))
	}
	if CrystalBallSearch(list, 23) != 11 {
		t.Error("BinarySearch failed. Expected 11, got ", CrystalBallSearch(list, 23))
	}
	if CrystalBallSearch(list, 4) != -1 {
		t.Error("BinarySearch failed. Expected -1, got ", CrystalBallSearch(list, 4))
	}
}

func TestCrystalBallSearchJump7(t *testing.T) {
	list := slice.GenerateSlice(45, 1, 2)
	// check indexes 0, 1, 6, 7, 23, 27, 42, 45

	if CrystalBallSearch(list, 1) != 0 {
		t.Error("BinarySearch failed. Expected 0, got ", CrystalBallSearch(list, 1))
	}
	if CrystalBallSearch(list, 3) != 1 {
		t.Error("BinarySearch failed. Expected 1, got ", CrystalBallSearch(list, 3))
	}
	if CrystalBallSearch(list, 13) != 6 {
		t.Error("BinarySearch failed. Expected 6, got ", CrystalBallSearch(list, 13))
	}
	if CrystalBallSearch(list, 15) != 7 {
		t.Error("BinarySearch failed. Expected 7, got ", CrystalBallSearch(list, 15))
	}
	if CrystalBallSearch(list, 47) != 23 {
		t.Error("BinarySearch failed. Expected 23, got ", CrystalBallSearch(list, 47))
	}
	if CrystalBallSearch(list, 55) != 27 {
		t.Error("BinarySearch failed. Expected 27, got ", CrystalBallSearch(list, 55))
	}
	if CrystalBallSearch(list, 85) != 42 {
		t.Error("BinarySearch failed. Expected 42, got ", CrystalBallSearch(list, 85))
	}
	if CrystalBallSearch(list, 89) != 44 {
		t.Error("BinarySearch failed. Expected 44, got ", CrystalBallSearch(list, 89))
	}

	if CrystalBallSearch(list, 99) != -1 {
		t.Error("BinarySearch failed. Expected -1, got ", CrystalBallSearch(list, 99))
	}
	if CrystalBallSearch(list, 90) != -1 {
		t.Error("BinarySearch failed. Expected -1, got ", CrystalBallSearch(list, 90))
	}
	if CrystalBallSearch(list, 4) != -1 {
		t.Error("BinarySearch failed. Expected -1, got ", CrystalBallSearch(list, 4))
	}

}

func TestCrystalBallSearchEmptySlice(t *testing.T) {
	var list []int

	if CrystalBallSearch(list, 3) != -1 {
		t.Error("BinarySearch failed. Expected -1, got ", CrystalBallSearch(list, 3))
	}
}

func TestCrystalBallSearch1ElementSlice(t *testing.T) {
	list := []int{1}

	if CrystalBallSearch(list, 3) != -1 {
		t.Error("BinarySearch failed. Expected -1, got ", CrystalBallSearch(list, 3))
	}
	if CrystalBallSearch(list, 1) != 0 {
		t.Error("BinarySearch failed. Expected 0, got ", CrystalBallSearch(list, 1))
	}
}
