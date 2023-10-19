package search

import (
	"math"
)

// CrystalBallSearch is a search algorithm that jump sqrt(n) times and then linear search
// Condition is having 2 times max comparison that element from slice is bigger than target
// This is not really crystal ball search because crystal ball shouldnt receive target, but this is similar algorithm
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
