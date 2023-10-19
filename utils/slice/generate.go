package slice

func GenerateSlice(len, start, step int) []int {
	s := make([]int, len)
	for i := range s {
		s[i] = step*i + start
	}
	return s
}
