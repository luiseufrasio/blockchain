package main

import "fmt"

func main() {
	myMap := make(map[int]string)

	fmt.Println(myMap)

	myMap[52] = "1st"
	myMap[23] = "2nd"

	fmt.Println(myMap)
}