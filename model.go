package main

import "errors"

type User struct {
	Id   int
	Name string
}

func CreateUser(id int, name string) (*User, error) {
	if !IsValidId(id) {
		return nil, errors.New("invalid id")
	}
	if !IsValidName(name) {
		return nil, nil
	}
	return &User{id, name}, nil
}
