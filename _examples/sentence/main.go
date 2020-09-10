package main

import (
	"fmt"
	"strings"
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
	//  I am a free man! I am not a number!
}

func f(str string) {
	order := 2
	text := " " + str
	// occurrences := map[string]int{}
	words := strings.Split(text, " ")
	pairs := []string{" "}

	for i := 0; i < len(words)-order; i++ {
		pairs = append(pairs, strings.Join(words[i:i+order], " "))
	}

	for i := range pairs {
		fmt.Println(pairs[i])
	}

}
