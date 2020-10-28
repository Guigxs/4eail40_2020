package coord

import (
	"fmt"
)

// ICoord is the coordinate interface.
// GetCoord return the index of the coord
type ICoord interface {
	fmt.Stringer
	GetCoord(int) int
}

// CartesianCoord is an ICoord implementation that use a X, Y axis
type CartesianCoord struct {
	X string
	Y int
}

func (c *CartesianCoord) String() string {
	return fmt.Sprintf("%s%d equivalent to [%d, %d]", c.X, c.Y, c.Y, c.GetAlphabetToint(c.X))
}

// GetCoord return the Line, Row indexes for the board.Table
func (c *CartesianCoord) GetCoord(id int) int {
	switch id {
	case 1:
		return c.GetAlphabetToint(c.X)
	case 0:
		return c.Y
	}
	return -1
}

// GetAlphabetToInt return the letter index in the alphabet
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
