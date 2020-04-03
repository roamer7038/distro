// distro パッケージは，確率分布に基づいた疑似乱数ジェネレータを提供します．
package distro

import (
	"fmt"
	"math"
	"math/rand"
)

// Zipf はZipf則に基づく確率分布を保持する構造体です．
type Zipf struct {
	r *rand.Rand
	a float64
	n int
	f []float64
	c []float64
}

// NewZipf はRand構造体と総コンテンツ数，偏りパラメータを引数に，Zipf構造体を生成します．
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

// Pdf は指定した順位のアクセス確率を返します．
func (z *Zipf) Pdf(rank int) (float64, error) {
	var err error

	if rank < 0 || z.n < rank {
		err = fmt.Errorf("Invalid parameter n=%d", z.n)
	}
	return z.f[rank-1], err
}

// Cdf は指定した順位までの累積確率を返します．
func (z *Zipf) Cdf(rank int) (float64, error) {
	var err error

	if rank < 0 || z.n < rank {
		err = fmt.Errorf("Invalid parameter n=%d", z.n)
	}
	return z.c[rank-1], err
}

// Uint64 はZipf則に基づいた疑似乱数を返します．
func (z *Zipf) Uint64() uint64 {
	var rank uint64
	r := z.r.Float64()

	for i := 1; i <= z.n; i++ {
		if r < z.c[i-1] {
			rank = uint64(i)
			break
		}
	}

	return rank
}
