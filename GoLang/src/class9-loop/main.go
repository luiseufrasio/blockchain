package main

func main() {

	i := 0
	for {
		i++
		println(i)

		if i > 10 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	array := []int{10, 100, 1000}
	for i, v := range array {
		println(i, v)
	}

	map1 := make(map[string]string)
	map1["1"] = "Learning"
	map1["2"] = "Go"
	map1["3"] = "Language"
	for key, value := range map1 {
		println(key, value)
	}

}