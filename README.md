# polo 

A Markov chain implementation in Go.

![Chain](./chain.png)

```bash
➜  margov git:(master) ✗ go run .
P(Sunny|Sunny) = 0.80
P(Rainy|Sunny) = 0.05
P(Cloudy|Sunny) = 0.15

P(Sunny|Rainy) = 0.20
P(Rainy|Rainy) = 0.60
P(Cloudy|Rainy) = 0.20

P(Sunny|Cloudy) = 0.20
P(Rainy|Cloudy) = 0.30
P(Cloudy|Cloudy) = 0.50


The probability of tomorrow being Sunny, given that today was Rainy is 0.2

The next state, given that today was Sunny is Sunny
```

## TODO

- [x] Basic Markov Chain
- [x] Fix `Chain.Next()`
- [x] Implement `cumsum` and implement better `rand.Sample` algorithm
- [ ] Drawing/Visualizing a graph of chain with probabilities
- [x] Higher Order Chains 

## References

- <http://cecas.clemson.edu/~ahoover/ece854/refs/Ramos-Intro-HMM.pdf>
- <https://www.davidsilver.uk/wp-content/uploads/2020/03/MDP.pdf>
- <https://www.wikiwand.com/en/Pseudo-random_number_sampling>

