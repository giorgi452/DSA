package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less []int
	var greater []int

	for _, i := range arr[1:] {
		if i <= pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}
	result := append(quicksort(less), pivot)
	result = append(result, quicksort(greater)...)

	return result
}

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 3, 7, 5, 99}))
}
