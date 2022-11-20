package main

import "fmt"

func main() {
	user := myUser()
	user.user_map["name"] = "José"

	fmt.Println(user)
	user.printName();
}

type User struct {
	user_map map[string]string
}

func myUser() *User {
	result := User{}
	result.user_map = map[string]string{}
	
	return &result
}

func (user *User) printName() {
	println(user.user_map["name"])
}