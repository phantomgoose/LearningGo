package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(output []byte) (int, error) {
	copy(output, []byte("A"))
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
