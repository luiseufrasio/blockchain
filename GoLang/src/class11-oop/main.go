package main

func main() {
	user := User{}
	user.name = "Jose"

	println(user.name)
}

type User struct {
	name string
}