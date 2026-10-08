package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

func main() {
	account := Account{Owner: "Sara", Balance: 70}
	fmt.Printf("%+v\n", account)
}
