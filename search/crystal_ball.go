package search

import (
	"math"
)

func CrystalBallSearch(list []int, target int) int {
	jump := int(math.Ceil(math.Sqrt(float64(len(list)))))
	var low, high int
	for i := 0; i < jump; i++ {
		current := (i + 1) * jump
		if current >= len(list) || list[current] > target {
			low = current - jump
			high = int(math.Min(float64(current), float64(len(list))))
			break
		}
		if list[current] == target {
			return current
		}
	}
	for i := low; i < high; i++ {
		if list[i] == target {
			return i
		}
	}
	return -1
}
