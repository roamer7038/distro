package distro

import (
	"fmt"
	"math"
	"math/rand"
)

type Gamma struct {
	r *rand.Rand
	k float64
	t float64
	f []float64
	n int
	c []float64
}

func NewGamma(r *rand.Rand, n int, k float64, thita float64) (*Gamma, error) {
	var err error

	if n <= 0 {
		err = fmt.Errorf("Invalid parameter n=%d", n)
	}

	var (
		frequency  []float64
		cumulative []float64
	)

	sum := cdf(n, k, thita)
	for i := 0; i < n; i++ {
		frequency = append(frequency, pdf(i+1, k, thita)/sum)

		if i == 0 {
			cumulative = append(cumulative, frequency[i])
		} else {
			cumulative = append(cumulative, cumulative[i-1]+frequency[i])
		}
	}

	return &Gamma{
		r: r,
		k: k,
		t: thita,
		f: frequency,
		n: n,
		c: cumulative,
	}, err
}

func pdf(x int, k float64, thita float64) float64 {
	return math.Pow(float64(x), k-1) / (math.Gamma(k) * math.Pow(thita, k)) * math.Exp(float64(-x)/thita)
}

func cdf(x int, k float64, thita float64) float64 {
	var value float64
	for i := 1; i <= x; i++ {
		value += pdf(i, k, thita)
	}
	return value
}

func (self *Gamma) Pdf(x int) (float64, error) {
	var err error

	if x < 0 || self.n < x {
		err = fmt.Errorf("Invalid parameter n=%d", self.n)
	}
	return self.f[x-1], err
}

func (self *Gamma) Cdf(x int) (float64, error) {
	var err error

	if x < 0 || self.n < x {
		err = fmt.Errorf("Invalid parameter n=%d", self.n)
	}

	return self.c[x-1], err
}

func (self *Gamma) Uint64() uint64 {
	var x uint64
	r := self.r.Float64()

	for i := 1; i <= self.n; i++ {
		if r < self.c[i-1] {
			x = uint64(i)
			break
		}
	}

	return x
}
