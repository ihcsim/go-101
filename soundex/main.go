package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

import "strings"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/testsuite", testSuite)
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}

func home(response http.ResponseWriter, request *http.Request) {
	if queries := request.URL.Query(); len(queries) > 0 {
		if names, ok := queries["name"]; ok {
			body := make([]string, 0)
			for _, name := range names {
				code := Soundex(name)
				body = append(body, fmt.Sprintf("%s => [%s]", name, code))
			}

			response.WriteHeader(http.StatusOK)
			response.Write([]byte(strings.Join(body, "\n")))
		}
	} else {
		defaultContent := `
From Wikipedia, https://en.wikipedia.org/wiki/Soundex

Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English. The goal is for homophones to be encoded to the same representation so that they can be matched despite minor differences in spelling.[1] The algorithm mainly encodes consonants; a vowel will not be encoded unless it is the first letter. Soundex is the most widely known of all phonetic algorithms (in part because it is a standard feature of popular database software such as DB2, PostgreSQL,[2] MySQL,[3] Ingres, MS SQL Server[4] and Oracle[5]) and is often used (incorrectly) as a synonym for "phonetic algorithm".[citation needed] Improvements to Soundex are the basis for many modern phonetic algorithms.[6].
`
		if _, err := response.Write([]byte(defaultContent)); err != nil {
			log.Printf("%v\n", err)
		}
	}
}

func testSuite(response http.ResponseWriter, request *http.Request) {
	if _, err := response.Write([]byte("Run tests")); err != nil {
		log.Printf("%v\n", err)
	}
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

	return strings.ToUpper(encoded.String())
}
