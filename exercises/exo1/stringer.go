package main

import (
	"fmt"
	"strconv"
)

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) String() string {
	return "Square of width " + strconv.Itoa(r.Width) + " and length " + strconv.Itoa(r.Length)
}

type Circle struct {
	Radius int
}

func (c Circle) String() string {
	return "Circle of radius " + strconv.Itoa(c.Radius)
}

type Shape interface {
	fmt.Stringer
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}

}
