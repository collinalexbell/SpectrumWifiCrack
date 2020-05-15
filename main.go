package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"
)

const bufferSize int = 1024 * 1024 * 5

func main() {
	fmt.Println("Spectrum Wifi Crack Expiriment")
	fmt.Println(time.Now())
	words := allWordsOfLength(5)
	fmt.Printf("enumurating psk of the shape <5 letter word><5 letter word><x : [0:1000]\n")
	fmt.Printf("Number of 5 letter words in dict: %d\n", len(words))
	enumerations := len(words) * len(words) * 1000
	fmt.Printf("Total enumerations to try: %d\n", enumerations)
	gigabytes := float64(enumerations*(5+5+(2 /*average bytes in x*/)+(1 /*newline*/))) / 1e9
	fmt.Printf("Password enumeration file in Gigabytes: %f\n", gigabytes)

	enumeratePotentialPasswords(words)
	fmt.Println(time.Now())
}

func enumeratePotentialPasswords(words [][]byte) {
	length := float64(len(words))
	buffer := make([]byte, bufferSize)
	ints := rangeAsBytes(1000)
	bufferIndex := 0
	var bufferLength int
	for index, first := range words {
		for _, second := range words {
			for _, x := range ints {
				bufferIndex = copyPsk(buffer, first, second, x, bufferIndex)

				if bufferIndex+15 >= bufferSize {
					bufferLength = bufferIndex
					saveBuffer(buffer, bufferLength)
					bufferIndex = 0
				}
			}
		}
		if index%50 == 0 {
			fmt.Printf("%f%% finished\n", float64(index)/length*100)
		}
	}
	bufferLength = bufferIndex
	saveBuffer(buffer, bufferLength)

}

func saveBuffer(buffer []byte, bufferLength int) {
	// This function can be used to either write the buffer to a disk or to give it to program like hashcat.
	// To be implemented later (when I have the disk space to actually save the entire file)
}

var newLine []byte = []byte("\n")
var lenN int

func copyPsk(buffer []byte, first []byte, second []byte, x []byte, bufferIndex int) int {
	copy(buffer[bufferIndex:bufferIndex+5], first)
	copy(buffer[bufferIndex+5:bufferIndex+10], second)
	lenN = len(x)
	copy(buffer[bufferIndex+10:bufferIndex+10+lenN], x)
	copy(buffer[bufferIndex+10+lenN:bufferIndex+10+lenN+1], newLine)
	return bufferIndex + 11 + lenN
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
