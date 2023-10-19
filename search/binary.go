package search

// BinarySearch performs a binary search on a sorted slice of ints.
// [lo,hi) is the half-open range of indices to search.
func BinarySearch(list []int, target int) int {
	var mid, guess, low int
	high := len(list)
	for low < high {
		mid = low + (high-low)/2
		guess = list[mid]
		if guess == target {
			return mid
		} else if guess < target {
			low = mid + 1
		} else if guess > target {
			high = mid
		}
	}
	return -1
}
