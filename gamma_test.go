package distro

import (
	"math/rand"
	"testing"
	"time"
)

func TestGamma(t *testing.T) {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	g, _ := NewGamma(rs, 1000, 0.475, 170.607)

	pdf, _ := g.Pdf(1)
	cdf, _ := g.Cdf(1)
	if pdf != cdf {
		t.Errorf("A value of Pdf(1) and Cdf(1) do not match.\n pdf: %v\ncdf: %v", pdf, cdf)
	}

	expect := 0.050099
	if round(pdf, 6) != expect {
		t.Errorf("A value of Pdf(1) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	pdf, _ = g.Pdf(1000)
	expect = 0.000382
	if round(pdf*100, 6) != expect {
		t.Errorf("A value of Pdf(1000) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	cdf, _ = g.Cdf(1000)
	expect = 1
	if round(cdf, 6) != expect {
		t.Errorf("A value of Cdf(1000) do not match.\n got: %v\nwant: %v", cdf, expect)
	}

	r := g.Uint64()
	if r < 0 && 1000 < r {
		t.Errorf("An inappropriate value is returned from Rand(). got: %v", r)
	}
}
