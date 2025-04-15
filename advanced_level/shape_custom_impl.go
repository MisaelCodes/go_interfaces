package advanced_level

import (
	"fmt"
	"math"
)

type Shape interface{
    Area() float64
}

type Circle struct{
    Radious float64
}

func (c *Circle) Area() float64{
    return math.Pi * math.Pow(c.Radious, 2)
}

type Rectangle struct{
    Width float64
    Height float64
}

func (r *Rectangle) Area() float64{
    return r.Width * r.Height
}

func PrintArea(s Shape){
    fmt.Println(s.Area())
}
