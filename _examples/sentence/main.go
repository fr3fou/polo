package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fr3fou/margov/margov"
)

type DM struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func main() {
	if len(os.Args) < 2 {
		panic("not enough args, provide path to fb inbox dir")
	}
	dir := os.Args[1]

	sentences := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		var d DM
		if err := json.NewDecoder(f).Decode(&d); err != nil {
			return err
		}

		for _, m := range d.Messages {
			if m.Type != "Generic" {
				continue
			}

			sentences = append(sentences, m.Content)
		}

		return f.Close()
	})
	if err != nil {
		panic(err)
	}

	order := 1
	occurrences := buildOccurrences(sentences, order)
	chain := createChain(occurrences, order)
	fmt.Println(predict(chain, "da", 5))
}

func buildOccurrences(sentences []string, order int) map[string]map[string]int {
	occurrences := map[string]map[string]int{}
	for _, str := range sentences {
		text := " " + str // Pad the beginning with empty string
		words := strings.Split(text, " ")

		// TODO: optimize this
		for i := 0; i < len(words)-order; i++ {
			pair := strings.Join(words[i:i+order], " ")

			if _, ok := occurrences[pair]; !ok {
				occurrences[pair] = map[string]int{}
			}

			occurrences[pair][words[i+order]]++
		}
	}
	return occurrences
}

func createChain(m map[string]map[string]int, order int) margov.Chain {
	chain := margov.New(order)

	// TODO: optimize this
	for pair, words := range m {
		total := float64(len(words))
		for word, occurrence := range words {
			chain.Set(word, float64(occurrence)/total, strings.Split(pair, " ")...)
		}
	}

	return chain
}

func predictRandom(c margov.Chain, n int) string {
	return ""
}

func predict(c margov.Chain, input string, n int) string {
	final := input + " "
	next := input

	for i := 0; i < n; i++ {
		next = c.Next(next)
		final += next + " "
	}

	return final
}
