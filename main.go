package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fr3fou/margov/margov"
)

const (
	Sunny  = "Sunny"
	Rainy  = "Rainy"
	Cloudy = "Cloudy"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func _main() {
	chain := margov.New()

	// Given that the current day is Sunny.
	chain.Set(Sunny, Sunny, 0.8)
	chain.Set(Rainy, Sunny, 0.05)
	chain.Set(Cloudy, Sunny, 0.15)

	// Given that the current day is Rainy.
	chain.Set(Sunny, Rainy, 0.2)
	chain.Set(Rainy, Rainy, 0.6)
	chain.Set(Cloudy, Rainy, 0.2)

	// Given that the current day is Rainy.
	chain.Set(Sunny, Cloudy, 0.2)
	chain.Set(Rainy, Cloudy, 0.3)
	chain.Set(Cloudy, Cloudy, 0.5)

	fmt.Println(chain)

	fmt.Println("The probability of tomorrow being Sunny, given that today was Rainy is", chain.Probability(Sunny, Rainy))

	fmt.Println()
	fmt.Println("The next state, given that today was Sunny is", chain.Next(Sunny))
}
