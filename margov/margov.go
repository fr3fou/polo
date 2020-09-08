package margov

import (
	"fmt"
	"math/rand"
	"strings"
)

// State is a string.
type State = string

// Chain is a Sequence of random states -> probabilities.
type Chain map[State]Probabilities

// Probabilities gives the probabilities of going to the next state given some state.
type Probabilities map[State]float64

// New is a constructor of Chain.
func New() Chain {
	return Chain{}
}

// Set sets the probability matrix of the current state to some next state.
func (c Chain) Set(next State, current State, probability float64) {
	// If the key state doesn't exist, initialize it
	if _, ok := c[current]; !ok {
		c[current] = Probabilities{}
	}

	c[current][next] = probability
}

func (c Chain) String() string {
	sb := strings.Builder{}

	for current, matrix := range c {
		for next, probability := range matrix {
			sb.WriteString(fmt.Sprintf("P(%s|%s) = %.02f\n", next, current, probability))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// Probability returns the probability of the next state happening given the current one.
func (c Chain) Probability(next State, current State) float64 {
	return c[current][next]
}

// Next gives the next state given the current state.
func (c Chain) Next(current State) State {
	probs := []float64{}
	states := []State{}

	for state, probability := range c[current] {
		probs = append(probs, probability)
		states = append(states, state)
	}

	sum := cumsum(probs)
	sample := rand.Float64()

	for index, val := range sum {
		if sample <= val {
			return states[index]
		}
	}

	return ""
}

func cumsum(p []float64) []float64 {
	sums := make([]float64, len(p))
	sums[0] = p[0]

	for i := 1; i < len(p); i++ {
		sums[i] = sums[i-1] + p[i]
	}

	return sums
}
