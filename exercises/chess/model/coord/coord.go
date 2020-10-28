package coord

import (
	"fmt"
)

type ICoord interface {
	fmt.Stringer
	GetCoord(int) int
}

type CartesianCoord struct {
	x byte
	y int
}

func (c *CartesianCoord) String() string {
	return fmt.Sprintf("[%s, %d] equivalent to [$d, %d]", c.x, c.y, c.GetAlphabetToint(c.x), c.y)
}

func (c *CartesianCoord) GetCoord(id int) int {
	switch id {
	case 0:
		return c.GetAlphabetToint(c.x)
	case 1:
		return c.y
	}
	return -1
}

func (c CartesianCoord) GetAlphabetToint(letter byte) int {
	alphabetMaj := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetMin := "abcdefghijklmnopqrstuvwxyz"

	for i := range alphabetMaj {
		if letter == alphabetMaj[i] || letter == alphabetMin[i] {
			return i
		}
	}
	return -1

}
