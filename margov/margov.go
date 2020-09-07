package margov

// State is a string.
type State = string

// Chain is a Sequence of random states -> probabilities.
type Chain map[State][][]float64

// New is a constructor of Chain
func New() Chain {
	return Chain{}
}
