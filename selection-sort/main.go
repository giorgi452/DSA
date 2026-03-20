package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(arr []int) []int {
	var newArr []int
	length := len(arr)

	for i := 0; i < length; i++ {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func main() {
	fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
}
