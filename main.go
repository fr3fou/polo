package main

import (
	"github.com/fr3fou/margov/margov"
)

const (
	Sunny  = "Sunny"
	Rainy  = "Rainy"
	Cloudy = "Cloudy"
)

func main() {

	chain := margov.New()
	chain.Set(Sunny, Sunny, 0.8)
	chain.Set(Sunny, Sunny, 0.8)
}
