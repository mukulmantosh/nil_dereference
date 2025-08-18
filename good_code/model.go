package main

import (
	"errors"
	"regexp"
	"sync/atomic"
)

var idCounter int64 = 0

type User struct {
	ID   int64
	Name string
}

func CreateUser(name string, email string) (*User, error) {
	if !IsValidEmail(email) {
		return nil, errors.New("the email is invalid")
	}
	if !IsValidName(name) {
		return nil, nil
	}
	return &User{ID: GenerateUniqueIntID(), Name: name}, nil
}

func IsValidName(name string) bool {
	return name != ""
}

func IsValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return true
	}
	return false
}

func GenerateUniqueIntID() int64 {
	return atomic.AddInt64(&idCounter, 1)
}
