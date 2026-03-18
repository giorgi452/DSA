package main

import (
	"fmt"
	"sort"
)

func binSearch(arr []string, item string) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	myArr := []string{"saxeli 1", "saxeli 2", "saxeli 3"}

	sort.Strings(myArr)
	foundIndex := binSearch(myArr, "saxeli 5")

	if foundIndex < 0 {
		fmt.Println("Not Found")
		return
	}
	fmt.Println(myArr[foundIndex])
}
