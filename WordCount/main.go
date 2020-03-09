package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) (wordMap map[string]int) {
	words := strings.Fields(s)
	wordMap = make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	return
}

func main() {
	wc.Test(wordCount)
}
