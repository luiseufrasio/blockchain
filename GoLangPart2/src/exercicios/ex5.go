package main

import "fmt"

func linearSearch(array []int, number int) {
	for i := 0; i < len(array); i++ {
		if array[i] == number {
			fmt.Printf("The int %d is in position %d", number, i)
			return
		}
	}
	fmt.Printf("The int %d was not found", number)
}

func main() {
	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	
	fmt.Println("Result for 1:")
	linearSearch(array, 1)
}