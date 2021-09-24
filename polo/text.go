package polo

import "strings"

const EndState = "EndEvent"

func NewFromText(order int, sentences []string) Chain {
	chain := New(order)
	occurrences := buildOccurrences(sentences, chain.Order)
	for prevStates, nextStates := range occurrences {
		total := float64(len(nextStates))
		for to, occurrence := range nextStates {
			prob := float64(occurrence) / total
			from := strings.Split(prevStates, " ")
			chain.Set(to, prob, from...)
		}
	}
	return chain
}

func buildOccurrences(states []State, order int) map[State]map[State]int {
	// map from previous states (looking back order amount of times)
	// to a map of the next state and the respective probability
	occurrences := map[State]map[State]int{}
	for _, str := range states {
		text := " " + str // Pad the beginning with empty string
		words := strings.Split(text, " ")

		var pair string
		for i := 0; i < len(words)-order; i++ {
			pair = strings.Join(words[i:i+order], " ")
			if _, ok := occurrences[pair]; !ok {
				occurrences[pair] = map[string]int{}
			}
			occurrences[pair][words[i+order]]++
		}
                if _, ok := occurrences[pair]; !ok {
		        occurrences[pair] = map[string]int{}
		}
		occurrences[pair][EndState]++
	}
	return occurrences
}
