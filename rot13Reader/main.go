package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(l byte) byte {
	isCapital := l >= 'A' && l <= 'Z'
	if !isCapital && (l < 'a' || l > 'z') {
		return l
	}

	l += 13
	if isCapital && l > 'Z' || !isCapital && l > 'z' {
		l -= 26
	}
	return l
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, letter := range b {
		b[i] = rot13(letter)
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
