package main

import "regexp"

func IsValidId(id int) bool {
	return id >= 0
}

func IsValidName(name string) bool {
	if len(name) <= 1 || len(name) >= 50 {
		return false
	}
	matched, _ := regexp.MatchString(`^[A-Za-z0-9]+$`, name)
	return matched
}

func (user *User) Validate() bool {
	return IsValidId(user.Id) && IsValidId(user.Id)
}
