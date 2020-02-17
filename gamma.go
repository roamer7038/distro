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
	var g Gamma

	if n <= 0 {
		err = fmt.Errorf("Invalid parameter n=%d", n)
	}

	var (
		frequency  []float64
		cumulative []float64
	)

	sum := g.cdf(n, k, thita)
	for i := 0; i < n; i++ {
		frequency = append(frequency, g.pdf(i+1, k, thita)/sum)

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

func (g Gamma) pdf(x int, k float64, thita float64) float64 {
	return math.Pow(float64(x), k-1) / (math.Gamma(k) * math.Pow(thita, k)) * math.Exp(float64(-x)/thita)
}

func (g Gamma) cdf(x int, k float64, thita float64) float64 {
	var value float64
	for i := 1; i <= x; i++ {
		value += g.pdf(i, k, thita)
	}
	return value
}

func (g *Gamma) Pdf(x int) (float64, error) {
	var err error

	if x < 0 || g.n < x {
		err = fmt.Errorf("Invalid parameter n=%d", g.n)
	}
	return g.f[x-1], err
}

func (g *Gamma) Cdf(x int) (float64, error) {
	var err error

	if x < 0 || g.n < x {
		err = fmt.Errorf("Invalid parameter n=%d", g.n)
	}

	return g.c[x-1], err
}

func (g *Gamma) Uint64() uint64 {
	var x uint64
	r := g.r.Float64()

	for i := 1; i <= g.n; i++ {
		if r < g.c[i-1] {
			x = uint64(i)
			break
		}
	}

	return x
}
