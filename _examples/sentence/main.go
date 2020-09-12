package main

import (
	"strings"

	"github.com/fr3fou/margov/margov"
)

func main() {
	str := "I am not a number! I am a free man! I am also a wow"
	order := 2
	occurrences := buildOccurrences(str, order)
	chain := createChain(occurrences, order)
}

func buildOccurrences(str string, order int) map[string]map[string]int {
	text := " " + str // Pad the beginning with empty string
	occurrences := map[string]map[string]int{}
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-order; i++ {
		pair := strings.Join(words[i:i+order], " ")

		if _, ok := occurrences[pair]; !ok {
			occurrences[pair] = map[string]int{}
		}

		occurrences[pair][words[i+order]]++
	}

	return occurrences
}

func createChain(m map[string]map[string]int, order int) margov.Chain {
	chain := margov.New(order)

	for pair, words := range m {
		total := float64(len(words))
		for word, occurrence := range words {
			chain.Set(word, float64(occurrence)/total, strings.Split(pair, " ")...)
		}
	}

	return chain
}

// func h(c margov.Chain, n int, str string) {
// 	order := 2
// 	text := " " + str
// 	words := strings.Split(text, " ")
// 	next := strings.Join(words[:order], " ")

// 	for i := 0; i < n; i++ {
// 		result := c.Next(next)
// 		fmt.Print(result, " ")
// 		next = strings.Join(words[i:i+order], " ")
// 	}
// }
