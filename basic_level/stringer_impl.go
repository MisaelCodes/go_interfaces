package basic_level

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string{
    return fmt.Sprintf("Title: %s, Author: %s", b.Title, b.Author)
}
