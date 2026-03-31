package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[(low+high)/2]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}

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
