package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fr3fou/polo/polo"
)

const (
	Sunny  = "Sunny"
	Rainy  = "Rainy"
	Cloudy = "Cloudy"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	chain := polo.New(1)

	// Given that the current day is Sunny.
	chain.Set(Sunny, 0.8, Sunny)
	chain.Set(Rainy, 0.05, Sunny)
	chain.Set(Cloudy, 0.15, Sunny)

	// Given that the current day is Rainy.
	chain.Set(Sunny, 0.2, Rainy)
	chain.Set(Rainy, 0.6, Rainy)
	chain.Set(Cloudy, 0.2, Rainy)

	// Given that the current day is Rainy.
	chain.Set(Sunny, 0.2, Cloudy)
	chain.Set(Rainy, 0.3, Cloudy)
	chain.Set(Cloudy, 0.5, Cloudy)

	fmt.Println(chain)

	fmt.Println("The probability of tomorrow being Sunny, given that today was Rainy is", chain.Probability(Sunny, Rainy))

	fmt.Println()
	fmt.Println("The next state, given that today was Sunny is", chain.Next(Sunny))
	fmt.Println(chain.Graph())
}
