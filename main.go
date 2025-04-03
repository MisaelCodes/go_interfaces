package main

import (
	"fmt"

	"github.com/MisaelCodes/go_interfaces/basic_level"
)

func main(){
    b := basic_level.Book{Title: "Go Programming", Author: "John Doe"}
    fmt.Println(b)
}
