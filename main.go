package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

const bufferSize int = 1024 * 1024 * 5

func main() {
	words := allWordsOfLength(5)
	enumeratePotentialPasswords(words)
}

func enumeratePotentialPasswords(words [][]byte) {
	for _, first := range words {
		for _, second := range words {
			for x := 0; x < 1000; x++ {
        fmt.Printf("%s%s%d\n", first, second, x)
			}
		}
	}
}

func isLowerLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func allWordsOfLength(length int) [][]byte {
	rv := make([][]byte, 0)
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic("couldn't open the dictionary")
	}
	for _, word := range strings.Split(string(content), "\n") {
		if len(word) == length && isLowerLetters(word) {
			rv = append(rv, []byte(word))
		}
	}
	return rv
}
