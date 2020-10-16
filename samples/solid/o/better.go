package o

import "math"

// Calculator -- BAD --

//type Rectangle struct {
//	Width  int
//	Height int
//}

func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

//type Circle struct {
//	Radius int
//}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

type Shape interface {
	Area() float64
}

func ComputeAreaBetter(shapes []Shape) (totalArea float64) {
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return
}
