package margov

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
func (c Chain) Set(current State, next State, probability float64) {
	// If the key state doesn't exist, initialize it
	if _, ok := c[current]; !ok {
		c[current] = Probabilities{}
	}

	c[current][next] = probability
}
