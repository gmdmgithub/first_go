package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	totalN   int = 49
	lotteryN int = 6
	series   int = 5
)

// Lotto - struct for lottery numbers
type Lotto struct {
	generated []int
	rng       *rand.Rand
}

// NewLotto - generate slice of Lotto struct of length series
func NewLotto() []Lotto {

	lottos := make([]Lotto, series)

	for i := 0; i < series; i++ {
		seed := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(seed)

		l := &Lotto{
			generated: make([]int, 0),
			rng:       rng,
		}
		lottos[i] = *l.generate()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}

	return lottos
}

func (lotto *Lotto) generate() *Lotto {
	generated := map[int]bool{}

	for {
		if len(generated) >= lotteryN {
			break
		}

		v := lotto.rng.Int()%totalN + 1

		if !generated[v] {
			generated[v] = true
		}
	}

	for k := range generated {
		lotto.generated = append(lotto.generated, k)
	}

	sort.Ints(lotto.generated)

	return lotto
}

func main() {
	lottos := NewLotto()
	for _, lotto := range lottos {
		fmt.Printf("%v\n", lotto.generated)
	}
}
