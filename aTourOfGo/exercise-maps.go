package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ss := strings.Split(s, " ")
	m := make(map[string]int)
	for _, w := range ss {
		m[w]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
