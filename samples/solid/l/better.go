package l
import "fmt"

type Areaer interface {
	Area() int
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

type Square struct {
	Size int
}

func (s Square) Area() int {
	return s.Size * s.Size
}

func test() {
	r := Rectangle{2,3}
	fmt.Print("area: " + string(r.Area()))
}