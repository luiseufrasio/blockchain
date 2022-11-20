package main

import "fmt"

func swap(array []int, i, j int) {
	k := array[i]
	array[i] = array[j]
	array[j] = k
}

func bubbleSort(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array) - 1; i++ {
		if array[i] > array[i+1] {
			swap(array, i, i+1)
		}
	}
	bubbleSort(array[:len(array)-1])
}

func main() {
	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unordered array: ", array)
	bubbleSort(array)
	fmt.Println("Ordered array: ", array)
}