package polo

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/emicklei/dot"
)

// State is a string.
type State = string

// Chain is a Sequence of random states -> probabilities.
type Chain struct {
	StateTransitions map[State]Probabilities
	Order            int
}

// Probabilities gives the probabilities of going to the next state given some state.
type Probabilities map[State]float64

// New is a constructor of Chain.
func New(order int) Chain {
	return Chain{
		Order:            order,
		StateTransitions: map[string]Probabilities{},
	}
}

// Set sets the probability matrix of the current state to some next state.
func (c Chain) Set(next State, probability float64, states ...State) {
	if len(states) != c.Order {
		panic("Wrong amount of states provided")
	}

	current := strings.Join(states, " ")

	// If the key state doesn't exist, initialize it
	if _, ok := c.StateTransitions[current]; !ok {
		c.StateTransitions[current] = Probabilities{}
	}

	c.StateTransitions[current][next] = probability
}

func (c Chain) String() string {
	sb := strings.Builder{}

	for current, matrix := range c.StateTransitions {
		for next, probability := range matrix {
			sb.WriteString(fmt.Sprintf("P('%s'|'%s') = %.02f\n", next, current, probability))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// Probability returns the probability of the next state happening given the current one.
func (c Chain) Probability(next State, current State) float64 {
	return c.StateTransitions[current][next]
}

// Next gives the next state given the current state.
func (c Chain) Next(current State) State {
	probs := []float64{}
	states := []State{}

	for state, probability := range c.StateTransitions[current] {
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

	return current
}

func (c Chain) Graph() string {
	g := dot.NewGraph(dot.Directed)
	for from, probabilities := range c.StateTransitions {
		fromNode := g.Node(from)
		for to, prob := range probabilities {
			toNode := g.Node(to)
			g.Edge(fromNode, toNode, fmt.Sprintf("%.2f", prob))
		}
	}
	return g.String()
}

func cumsum(p []float64) []float64 {
	sums := make([]float64, len(p))
	sum := 0.0

	for i, p := range p {
		sum += p
		sums[i] = sum
	}

	return sums
}
