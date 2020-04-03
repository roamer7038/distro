# distro

Implement a pseudo-random number generator that follows an arbitrary probability distribution.

## Example

The following example implements a pseudo-random number generator that follows the Zipf distribution.
```
package main

import (
	"fmt"
	"github.com/roamer7038/distro"
	"math/rand"
	"time"
)

func main() {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	z, _ := distro.NewZipf(rs, 1000, 0.693)

	for i := 0; i < 1000; i++ {
		fmt.Println(z.Uint64())
	}
}
```

## API

Check the documentation using GoDoc (Written in Japanese).

```
godoc -http=:8000
```

### func NewZipf(r \*Rand, n int, alpha float64) (\*Zipf, error)
Generate a probability distribution that follows the Zipf distribution.
- `n` : the number of elements.
- `alpha` : the value of the exponent characterizing the distribution.

### func (z \*Zipf) Pdf(x int) (float64, error)
the probability density function, which describes a probability taking of the specified rank.

### func (z \*Zipf) Cdf(x int) (float64, error)
the cumulative distribution function, which describes a probability up to specified order.

### func (z \*Zipf) Uint64() uint64
returns a pseudo-random 64-bit value according to Zipf distribution.

### func NewGamma(r \*Rand, n int, k float64, thita float64) (\*Gamma, error)
Generate a probability distribution that follows the Gamma distribution. 
What is important is the values obtained by this method are normalized.
- `k` : a shape parameter.
- `thita`: a scale parameter θ.

### func (g \*Gamma) Pdf(x int) (float64, error)

### func (g \*Gamma) Cdf(x int) (float64, error)

### func (g \*Gamma) Uint64() uint64

