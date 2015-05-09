package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(output []byte) (int, error) {
	reader.r.Read(output)
	for i := 0; i < len(output); i++ {
		if output[i] == ' ' || output[i] == '!' {
			continue
		}

		if letter := Rot13(output[i]); (IsALowerCaseLetter(letter)) || (IsAnUpperCaseLetter(letter)) {
			output[i] = letter
		} else if ExceededLowerCaseLettersRange(letter) {
			output[i] = 65 + letter - 90 - 1
		} else if ExceededUpperCaseLettersRange(letter) {
			output[i] = 97 + letter - 122 - 1
		}
	}

	return len(output), io.EOF
}

func Rot13(input uint8) (newLetter uint8) {
	return input + 13
}

func IsALowerCaseLetter(input uint8) bool {
	return input >= 97 && input <= 122
}

func IsAnUpperCaseLetter(input uint8) bool {
	return input >= 65 && input <= 90
}

func ExceededLowerCaseLettersRange(input uint8) bool {
	return input > 90 && input < 97
}

func ExceededUpperCaseLettersRange(input uint8) bool {
	return input > 122
}
