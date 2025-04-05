package basic_level

import "fmt"

type Speaker interface {
	Speak()
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("miu i'm gonna conquer the world")
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Guau love u")
}

func MakeSound(s Speaker) {
	s.Speak()
}
