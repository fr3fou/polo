# margov
A Markov chain implementation in Go.

![Chain](./chain.png)

```bash
➜  margov git:(master) ✗ go run .
P(Cloudy|Sunny) = 1.15
P(Sunny|Sunny) = 0.80
P(Rainy|Sunny) = 0.05

P(Cloudy|Rainy) = 0.20
P(Sunny|Rainy) = 0.20
P(Rainy|Rainy) = 0.60

P(Sunny|Cloudy) = 0.20
P(Rainy|Cloudy) = 0.30
P(Cloudy|Cloudy) = 0.50


The probability of tomorrow being Sunny, given that today was Rainy is 0.2

The next state, given that today was Sunny is Cloudy
```

## References

- <http://cecas.clemson.edu/~ahoover/ece854/refs/Ramos-Intro-HMM.pdf>
- <https://www.davidsilver.uk/wp-content/uploads/2020/03/MDP.pdf>

