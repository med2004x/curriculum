package main

import "fmt"

type Stringer interface {
	String() string
}

type User struct {
	Name string
}

func (u User) String() string {
	// TODO: define the representation contract.
	return u.Name
}

func describe(s Stringer) string {
	return s.String()
}

func main() {
	fmt.Println(describe(User{Name: "Ada"}))
}
