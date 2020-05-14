package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	fmt.Println("Spectrum Wifi Crack Expiriment")
	fmt.Println(time.Now())
	words := allWordsOfLength(5)
	combos := 0
	length := float64(len(words))
	for index := range words {
		for range words {
			for x := 0; x < 1000; x++ {
				combos++
			}
		}
		if index%50 == 0 {
			fmt.Printf("%f%% finished\n", float64(index)/length*100)
		}
	}
	fmt.Printf("num combos: %v\n", combos)
	fmt.Println(time.Now())
}

func allWordsOfLength(length int) []string {
	rv := make([]string, 0)
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic("couldn't open the dictionary")
	}
	for _, word := range strings.Split(string(content), "\n") {
		if len(word) == length {
			rv = append(rv, word)
		}
	}
	return rv
}
