package main

import "sync/atomic"

var idCounter int64 = 0

type User struct {
	ID   int64
	Name string
}

func CreateUser(name string) *User {
	if !validName(name) {
		// ❌ Returns nil without signaling failure
		return nil
	}
	return &User{ID: GenerateUniqueIntID(), Name: name}
}

func validName(name string) bool {
	return name != ""
}

func GenerateUniqueIntID() int64 {
	return atomic.AddInt64(&idCounter, 1)
}
