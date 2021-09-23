package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chzyer/readline"
	"github.com/fr3fou/polo/polo"
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

	chain := polo.NewFromText(1, sentences)
	fmt.Println("Press enter for the next generated message")
	fmt.Println("	You can also enter a starting word")
	fmt.Println("	Type 'quit' to quit")
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	in := ""
	for {
		in, err = rl.ReadlineWithDefault(in)
		if err != nil {
			fmt.Println(err)
			return
		}
		if in == "quit" {
			return
		}
		fmt.Print("< ")
		if in == "" {
			fmt.Println(chain.NextUntilEnd(chain.RandomState()))
		} else {
			fmt.Println(chain.NextUntilEnd(in))
		}
	}
}
