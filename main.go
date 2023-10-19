package main

import "DsaPrimeagen/search"

func main() {
	list := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	i := search.CrystalBallSearch(list, 239)
	println(i)
}
