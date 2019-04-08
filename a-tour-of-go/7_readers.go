package main

import "golang.org/x/tour/reader"

type AReader struct{}

func (r AReader) Read(bytes []byte) (int, error) {
	for item := range bytes {
		bytes[item] = 65
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(AReader{})
}
