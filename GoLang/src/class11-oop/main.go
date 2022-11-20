package main

func main() {
	user := User{"Maria"}
	//user.name = "Jose"

	println(user.name)
}

type User struct {
	name string
}