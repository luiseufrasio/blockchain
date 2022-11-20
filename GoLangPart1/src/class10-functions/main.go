package main

// por referência
func main() {
	message := "Hey there!"
	hello(&message)
	println(message)

	println(sum(1, 9, 10, 8))
}

func hello(message *string) {
	println(*message)

	*message = "Hey UPDATED!"
}

func sum(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}

/* por valor
func main() {
	message := "Hey there!"
	hello(message)
}

func hello(message string) {
	println(message)
}
*/