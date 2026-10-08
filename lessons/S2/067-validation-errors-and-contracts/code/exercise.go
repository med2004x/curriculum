package main

import "fmt"

type CreateUser struct{ Name string }

func (u CreateUser) Validate() error {
	// TODO: define stable validation rules and error categories.
	if u.Name == "" {
		return fmt.Errorf("name required")
	}
	return nil
}

func main() { fmt.Println((CreateUser{Name: "Ada"}).Validate()) }
