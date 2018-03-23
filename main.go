package main

import "fmt"

import "github.com/instrument/go-zenefits/zenefits"

func main() {
	fmt.Println("hello world")
	foo := &zenefits.PeopleQueryParams{Limit: 100}
	fmt.Println(foo)
}
