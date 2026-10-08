package main

import "fmt"

type User struct {
	ID   int
	Role string
}

func canAccess(u User, resourceOwner int) bool {
	// TODO: separate authentication from authorization and enforce least privilege.
	return u.Role == "admin" || u.ID == resourceOwner
}

func main() { fmt.Println(canAccess(User{ID: 1, Role: "user"}, 1)) }
