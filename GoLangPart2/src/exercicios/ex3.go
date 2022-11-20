package main

import "fmt"

func merge(arr1 []int, arr2 []int) []int {
	resultArray := make([]int, len(arr1) + len(arr2))
	i := 0
	j:=0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			resultArray[i+j] = arr1[i]
			i++
		} else {
			resultArray[i+j] = arr2[j]
			j++
		}
	}
	for i < len(arr1) {
		resultArray[i+j] = arr1[i]
		i++
	}
	for j < len(arr2) {
		resultArray[i+j] = arr2[j]
		j++
	}
	return resultArray
}

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	middle := len(array) / 2
	part1 := mergeSort(array[:middle])
	part2 := mergeSort(array[middle:])
	return merge(part1, part2)
}

func main() {
	array := []int{1, 6, 2, 4, 9, 10, 5, 3, 7, 8}
	fmt.Println("Unordered array: ", array)
	fmt.Println("Ordered array: ", mergeSort(array))
}