# distro

[ja](./README.ja.md) | en

The package `distro` implements a pseudo-random number generator based on an arbitrary probability distribution. This package provides the following features:

- Generates a probability distribution based on arguments.
- Provides probability density functions (PDF) and cumulative distribution functions (CDF).
- Returns randomly the index value based on the probability distribution.

## Usage

### Installation

```
go get github.com/roamer7038/distro
```

### Example

In this example, we implement a pseudo-random number generator based on [Zipf's law](https://en.wikipedia.org/wiki/Zipf%27s_law). The program outputs the rankings randomly. 

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

To declare a random number generator, do the following. The argument `rs` is a pseudo-random number generator declared in the `rand` package, `n` is the number of elements and `alpha` is the value that characterizes the distribution.

```
z, err := distro.NewZipf(rs, n, alpha)
```

The probability density function, which describes a probability taking of the specified rank.

```
p, err := z.Pdf(rank)
```

The cumulative distribution function, which describes a probability up to specified order.

```
c, err := z.Cdf(rank)
```

Returns a pseudo-random 64-bit value based on Zipf's distribution.

```
k := z.Uint64()
```

## LICENSE

[MIT](./LICENSE)
