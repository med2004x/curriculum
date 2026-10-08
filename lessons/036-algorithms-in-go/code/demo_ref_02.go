package main

import "fmt"

type Student struct {
	Name   string
	Scores []int
}

func main() {
	students := []Student{
		{Name: "Sara", Scores: []int{70, 80, 60}},
		{Name: "Adam", Scores: []int{45, 50, 40}},
	}

	for _, student := range students {
		fmt.Println(student.Name, student.Scores)
	}
}
