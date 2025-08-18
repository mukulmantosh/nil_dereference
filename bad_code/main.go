package main

import "fmt"

func RegisterUser(name string) {
	user := CreateUser(name)
	fmt.Printf("Name: %s, ID: %d\n", user.Name, user.ID)
}

func main() {
	RegisterUser("")
}
