package main

import "fmt"

func findMinIndex(arr []int) int {
	minValue := arr[0]
	minIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}

	return minIndex
}

func selectionSort(arr []int) []int {
	size := len(arr)
	result := make([]int, size)
	for i := 0; i < size; i++ {
		minIndex := findMinIndex(arr)
		result[i] = arr[minIndex]
		arr = append(arr[:minIndex], arr[minIndex+1:]...)
	}
	return result
}

func main() {
	fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
}
