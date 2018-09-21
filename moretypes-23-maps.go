/*
https://tour.golang.org/moretypes/23
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)
	for _, v := range strings.Fields(s) {
		wordCountMap[v]++
	}
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
