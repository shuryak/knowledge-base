package main

import "fmt"

func quickSort(list []int) []int {
	// Базовый случай:
	if len(list) < 2 {
		return list
	}

	pivot := list[0] // опорный элемент

	// Строим подмассивы всех элементов, меньших/больших опорного
	var less = []int{}
	var greater = []int{}
	for _, num := range list[1:] {
		if pivot > num {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	// Рекурсивный случай:
	less = append(quickSort(less), pivot)
	greater = quickSort(greater)

	// Собираем результирующий массив
	return append(less, greater...)
}

func main() {
	fmt.Println(quickSort([]int{10, 5, 2, 3}))
}
