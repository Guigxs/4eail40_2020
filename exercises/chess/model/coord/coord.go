package coord

import (
	"fmt"
)

type ICoord interface {
	fmt.Stringer
	GetCoord(int) int
}

type CartesianCoord struct {
	X string
	Y int
}

func (c *CartesianCoord) String() string {
	return fmt.Sprintf("%s%d equivalent to [%d, %d]", c.X, c.Y, c.Y, c.GetAlphabetToint(c.X))
}

func (c *CartesianCoord) GetCoord(id int) int {
	switch id {
	case 1:
		return c.GetAlphabetToint(c.X)
	case 0:
		return c.Y
	}
	return -1
}

func (c CartesianCoord) GetAlphabetToint(letter string) int {
	alphabetMaj := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetMin := "abcdefghijklmnopqrstuvwxyz"

	for i := range alphabetMaj {
		if letter == string(alphabetMaj[i]) || letter == string(alphabetMin[i]) {
			return i
		}
	}
	return -1

}
