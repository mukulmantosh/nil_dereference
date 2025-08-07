package main

import (
	"errors"
)

var globalCounter int64

type User struct {
	Id   int
	Name string
}

type Ctx struct {
	DebugEnabled bool
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

func (user *User) Copy(ctx *Ctx) *User {
	if ctx.DebugEnabled {
		println("copy user...")
	}
	if !user.Validate() || user == nil {
		return nil
	}
	return &User{
		Id:   user.Id,
		Name: user.Name,
	}
}
