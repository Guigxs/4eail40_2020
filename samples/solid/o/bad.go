package o

import "math"

// Calculator -- BAD --

type Rectangle struct {
	Width  int
	Height int
}

type Circle struct {
	Radius int
}

func ComputeArea(shapes []interface{}) (totalArea float64) {
	for _, shape := range shapes {
		switch s := shape.(type) { // <-- switch on a type should be a warning
		case Rectangle:
			totalArea += float64(s.Width * s.Height)
		case Circle:
			totalArea += float64(s.Radius*s.Radius) * math.Pi
		}
	}
	return
}
