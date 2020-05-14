package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"
)

func main() {
	fmt.Println("Spectrum Wifi Crack Expiriment")
	fmt.Println(time.Now())
	words := allWordsOfLength(5)
	combos := 0
	//bufferIndex := 0
	length := float64(len(words))
	fmt.Printf("num words %d\n", int(length))
	const bufferSize int = 1024 * 1024 * 5
	//newLine := []byte("\n")
	//var buffer [bufferSize]byte

	ints := rangeAsBytes(1000)

	//var lenN int
	for index, _ := range words { //first
		for range words { //second
			for range ints { //x
				combos++
				/*
					copy(buffer[bufferIndex:bufferIndex+5], first)
					copy(buffer[bufferIndex+5:bufferIndex+10], second)
					lenN = len(x)
					copy(buffer[bufferIndex+10:bufferIndex+10+lenN], x)
					copy(buffer[bufferIndex+10+lenN:bufferIndex+10+lenN+1], newLine)
					bufferIndex = bufferIndex + 11 + lenN
					if bufferIndex+15 >= bufferSize {
						bufferIndex = 0
					}
				*/
			}
		}
		if index%50 == 0 {
			fmt.Printf("%f%% finished\n", float64(index)/length*100)
		}
	}

	fmt.Printf("num combos: %v\n", combos)
	fmt.Println(time.Now())
}

func rangeAsBytes(n int) [][]byte {
	rv := make([][]byte, 0)
	for x := 0; x < n; x++ {
		rv = append(rv, []byte(fmt.Sprintf("%d", x)))
	}
	return rv
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
