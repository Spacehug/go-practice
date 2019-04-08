package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := range b {
		if b[i] >= 0x41 && b[i] <= 0x5a {
			b[i] = byte((int(b[i])-0x41+13)%26 + 0x41)
		} else if b[i] >= 0x61 && b[i] <= 0x7a {
			b[i] = byte((int(b[i])-0x61+13)%26 + 0x61)
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
