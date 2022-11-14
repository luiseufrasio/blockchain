package main

import "fmt"

func main() {
	lista3 := []float32{7., 14., 49.}
	slice := lista3[:2]
	slice = append(slice, 21.)
	fmt.Println(slice)

	lista4 := make([]float32, 100)
	lista4[1] = 125
	fmt.Println(lista4)
	fmt.Println(len(lista4))

	/*
	lista2 := []int{6, 12, 36}
	fmt.Println(lista2)
	print(len(lista2))

	lista := [...]int{5, 25, 75}
	lista := [3]int{5, 25, 75}
	lista[0] = 5
	lista[1] = 25
	lista[2] = 75
	fmt.Println(lista)
	*/
}