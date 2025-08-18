package main

import (
	"fmt"
)

func RegisterUser(name string, email string) (*User, error) {
	user, err := CreateUser(name, email)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
	return user, nil
}

func main() {
	user, err := RegisterUser("", "mukul@goland.com")
	if err != nil {
		return
	}
	fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
}
