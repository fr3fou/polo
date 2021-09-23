package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chzyer/readline"
	"github.com/fr3fou/polo/polo"
)

type DM struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Author struct {
		Name string `json:"name"`
	} `json:"author"`
	Content string `json:"content"`
	ID      string `json:"id"`
	Type    string `json:"type"`
}

func main() {
	if len(os.Args) < 2 {
		panic("not enough args, provide path to discord dm json")
	}
	file := os.Args[1]
	mode := "repl"
	if len(os.Args) == 3 {
		mode = os.Args[2]
		if !(mode == "repl" || mode == "graph") {
			panic("mode must be repl or graph")
		}
	}
	messages := []string{}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	var d DM
	if err := json.NewDecoder(f).Decode(&d); err != nil {
		panic(err)
	}
	for _, message := range d.Messages {
		messages = append(messages, message.Content)
	}

	order := 1
	chain := polo.NewFromText(order, messages)
	if mode == "graph" {
		fmt.Println(chain.Graph())
		return
	}

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
