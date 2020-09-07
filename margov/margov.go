package margov

// State is a string.
type State = string

// Chain is a Sequence of random states -> probabilities.
type Chain map[State]Probabilities

type Probabilities map[State]float64

// New is a constructor of Chain
func New() Chain {
	return Chain{}
}

func (c Chain) Set(state State, otherState State, probability float64) {
	// If the key state doesn't exist, initialize it
	if _, ok := c[state]; !ok {
		c[state] = Probabilities{}
	}

	c[state][otherState] = probability
}
