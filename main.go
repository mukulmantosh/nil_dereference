package main

import (
	"fmt"
)

func SafeCreateUser(id int, name string) (*User, error) {
	user, err := CreateUser(id, name)
	if err != nil {
		return nil, fmt.Errorf("create user failed: %v", err)
	}
	fmt.Printf("Created user: %s, %d", user.Name, user.Id)
	return user, nil
}
