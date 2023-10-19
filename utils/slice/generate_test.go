package slice

import "testing"

func TestGenerateSlice(t *testing.T) {
	list := GenerateSlice(10, 1, 2)
	if len(list) != 10 || list[0] != 1 || list[9] != 19 {
		t.Error("GenerateSlice failed")
	}
}
