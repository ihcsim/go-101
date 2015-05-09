package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(output []byte) (int, error) {
	for i := 0; i < len(output); i++ {
		output[i] = byte('A')
	}
	return len(output), nil
}

func main() {
	reader.Validate(MyReader{})
}
