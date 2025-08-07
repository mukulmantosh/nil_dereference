package main

import (
	"fmt"
	"log"
)

func SafeCreateUser(id int, name string) (*User, error) {
	user, err := CreateUser(id, name)
	if err != nil {
		return nil, fmt.Errorf("create user failed: %v", err)
	}
	return user, nil
}

func SafeCreateUserDeeper(id int, name string) {
	user, err := SafeCreateUser(id, name)
	if err != nil {
		return
	}
	fmt.Printf("Name: %s, ID: %d", user.Name, user.Id)
}

// Analyzing function arguments
func example(user *User) {
	var ctx *Ctx
	if user == nil {
		log.Printf("user is nil")
	}
	user.Copy(ctx)
}

func main() {
	SafeCreateUserDeeper(1, "Mukul")
	example(&User{1, "Mukul"})
}
