package distro

import (
	"fmt"
	"math"
	"math/rand"
)

type Zipf struct {
	r *rand.Rand
	a float64
	n int
	f []float64
	c []float64
}

func NewZipf(r *rand.Rand, n int, alpha float64) (*Zipf, error) {
	var err error

	if n <= 0 {
		err = fmt.Errorf("Invalid parameter n=%d", n)
	}

	var (
		frequency   []float64
		numerator   float64
		denominator float64
		cumulative  []float64
	)

	for i := 1; i <= n; i++ {
		denominator += 1.0 / math.Pow(float64(i), alpha)
	}

	for i := 1; i <= n; i++ {
		numerator = 1.0 / math.Pow(float64(i), alpha)
		frequency = append(frequency, numerator/denominator)

		if i == 1 {
			cumulative = append(cumulative, frequency[i-1])
		} else {
			cumulative = append(cumulative, cumulative[i-2]+frequency[i-1])
		}
	}

	return &Zipf{
		r: r,
		a: alpha,
		n: n,
		f: frequency,
		c: cumulative,
	}, err
}

func (self *Zipf) Pdf(rank int) (float64, error) {
	var err error

	if rank < 0 || self.n < rank {
		err = fmt.Errorf("Invalid parameter n=%d", self.n)
	}
	return self.f[rank-1], err
}

func (self *Zipf) Cdf(rank int) (float64, error) {
	var err error

	if rank < 0 || self.n < rank {
		err = fmt.Errorf("Invalid parameter n=%d", self.n)
	}
	return self.c[rank-1], err
}

func (self *Zipf) Uint64() uint64 {
	var rank uint64
	r := self.r.Float64()

	for i := 1; i <= self.n; i++ {
		if r < self.c[i-1] {
			rank = uint64(i)
			break
		}
	}

	return rank
}
