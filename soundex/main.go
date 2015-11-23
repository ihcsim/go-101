package main

import "bytes"

import "strings"

func main() {

}

const codeLen = 4

var codes = map[string]string{
	"a": "",
	"b": "1",
	"c": "2",
	"d": "3",
	"e": "",
	"f": "1",
	"g": "2",
	"h": "",
	"i": "",
	"j": "2",
	"k": "2",
	"l": "4",
	"m": "5",
	"n": "5",
	"o": "",
	"p": "1",
	"q": "2",
	"r": "6",
	"s": "2",
	"t": "3",
	"u": "",
	"v": "1",
	"w": "",
	"x": "2",
	"y": "",
	"z": "2",
}

func Soundex(s string) string {
	var encoded bytes.Buffer
	encoded.WriteByte(s[0])

	for i := 1; i < len(s); i++ {
		if encoded.Len() == codeLen {
			break
		}

		previous, current := strings.ToLower(string(s[i-1])), strings.ToLower(string(s[i]))

		var next string
		if i+1 < len(s) {
			next = strings.ToLower(string(s[i+1]))
		}

		if (current == "h" || current == "w") && (codes[previous] == codes[next]) {
			i = i + 1
			continue
		}

		if c, ok := codes[current]; ok && len(c) > 0 {
			encoded.WriteByte(c[0])
		}

		if codes[current] == codes[next] {
			i = i + 1
			continue
		}
	}

	if encoded.Len() < codeLen {
		padding := strings.Repeat("0", codeLen-encoded.Len())
		encoded.WriteString(padding)
	}

	return encoded.String()
}
