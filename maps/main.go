package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func main() {
	wc.Test(WordCount)
}

func WordCount(input string) map[string]int {
	result := make(map[string]int)

	splits := strings.Fields(input)
	for i := 0; i < len(splits); i++ {
		count, exists := result[splits[i]]
		if exists {
			result[splits[i]] = count + 1
		} else {
			result[splits[i]] = 1
		}
	}

	return result
}
