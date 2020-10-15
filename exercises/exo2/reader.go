package main

import (
	// "fmt"
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (s spaceEraser) Read(b []byte) (n int, err error){
	n, err = s.r.Read(b)

	new_b := []byte{}
	for _, i := range(b){
		if i != 32{
			new_b = append(new_b, i)
		}
	}
	copy(b, new_b)
	return n, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
