package distro

import (
	"fmt"
	"math"
	"math/rand"
)

// Gamma はガンマ分布に基づく確率分布を保持する構造体です．
type Gamma struct {
	r *rand.Rand
	k float64
	t float64
	f []float64
	n int
	c []float64
}

// NewGamma はRand構造体と総コンテンツ数，形状記憶母数kと尺度母数θを引数に，Gamma構造体を生成します．
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

	if math.IsNaN(cumulative[n-1]) {
		err = fmt.Errorf("Normalization failed because the value is too large.")
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

// pdf は確率密度関数を提供します．
func (g Gamma) pdf(x int, k float64, thita float64) float64 {
	return math.Pow(float64(x), k-1) / (math.Gamma(k) * math.Pow(thita, k)) * math.Exp(float64(-x)/thita)
}

// cdf は累積確率密度関数を提供します．
func (g Gamma) cdf(x int, k float64, thita float64) float64 {
	var value float64
	for i := 1; i <= x; i++ {
		value += g.pdf(i, k, thita)
	}
	return value
}

// Pdf は指定した順位のアクセス確率を返します．
// この関数が提供する確率は正規化されています．
func (g *Gamma) Pdf(x int) (float64, error) {
	var err error

	if x < 0 || g.n < x {
		err = fmt.Errorf("Invalid parameter n=%d", g.n)
	}
	return g.f[x-1], err
}

// Cdf は指定した順位までの累積確率を返します．
// この関数が提供する確率は正規化されています．
func (g *Gamma) Cdf(x int) (float64, error) {
	var err error

	if x < 0 || g.n < x {
		err = fmt.Errorf("Invalid parameter n=%d", g.n)
	}

	return g.c[x-1], err
}

// Uint64 はガンマ分布に基づいた疑似乱数を提供します．
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
