package main

import (
	"fmt"

	"github.com/MisaelCodes/go_interfaces/basic_level"
)

func main() {
	b := basic_level.Book{Title: "Go Programming", Author: "John Doe"}
	fmt.Println(b)

	res, err := basic_level.Divide(10.0, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %v\n", res)
	}
	res, err = basic_level.Divide(10.0, 0)
	if err != nil {
		fmt.Println(err)
	}
}
