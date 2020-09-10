package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

/*
	" ":          {"I"},
	" I":         {"am"},
	"I am":       {"a", "not"},
	"a free":     {"man!"},
	"am a":       {"free"},
	"am not":     {"a"},
	"a number!":  {"I"},
	"number! I":  {"am"},
	"not a":      {"number!"},
*/
func main() {
	// chain := margov.New(2)

	f("I am not a number! I am a free man!")
	//  I am not a number! I am a free man!
}

func f(str string) {
	order := 2
	text := " " + str
	occurrences := map[string]map[string]int{}
	words := strings.Split(text, " ")
	// pairs := []string{" "}

	for i := 0; i < len(words)-order; i++ {
		pair := strings.Join(words[i:i+order], " ")

		if _, ok := occurrences[pair]; !ok {
			occurrences[pair] = map[string]int{}
		}

		occurrences[pair][words[i+order]]++
	}

	spew.Dump(occurrences)
}
