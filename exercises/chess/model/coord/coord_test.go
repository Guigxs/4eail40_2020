package coord

import (
	"testing"
)

func TestGetAlphabetToInt(t *testing.T) {
	tests := map[string]int{
		"A": 0,
		"E": 4,
		"z": 25,
		"$": -1,
	}
	for letter, want := range tests {
		t.Run(letter, func(t *testing.T){
			got := GetAlphabetToInt(letter)
			if got != want {
				t.Errorf("Error: '%s' = %d sould be %d", letter, got, want)
			}
		})
	}
}

func TestGetCoord(t *testing.T) {
	t.Run("A2", func (t *testing.T){
		a := CartesianCoord{X: "A", Y: 2}
		if a.GetCoord(0) != 2 || a.GetCoord(1) != 0 {
			t.Errorf("Error: GetCoord() sould return 2 and 0")
		}
	})
	t.Run("err", func(t *testing.T){
		a := CartesianCoord{X: "A", Y: 2}
		if got := a.GetCoord(2); got != -1{
			t.Errorf("Error: expected -1 but got %d", got)
		}
	})
}

func TestCartesianCoord_String(t *testing.T) {
	tests := []struct {
		name string
		c    CartesianCoord
		want string
	}{
		{
			"A0",
			CartesianCoord{X:"A", Y:0},
			"A0 equivalent to [0, 0]",
		},
		{
			"A2",
			CartesianCoord{X:"A", Y:2},
			"A2 equivalent to [2, 0]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("CartesianCoord.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
