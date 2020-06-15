package distro

import (
	"math/rand"
	"testing"
	"time"
)

func TestGamma(t *testing.T) {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	g, err := NewGamma(rs, 1000, 0.475, 170.607)
	catchError(t, err)

	pdf, err := g.Pdf(1)
	catchError(t, err)
	cdf, err := g.Cdf(1)
	catchError(t, err)
	if pdf != cdf {
		t.Errorf("A value of Pdf(1) and Cdf(1) do not match.\n pdf: %v\ncdf: %v", pdf, cdf)
	}

	expect := 0.050099
	if round(pdf, 6) != expect {
		t.Errorf("A value of Pdf(1) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	pdf, err = g.Pdf(1000)
	catchError(t, err)
	expect = 0.000382
	if round(pdf*100, 6) != expect {
		t.Errorf("A value of Pdf(1000) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	cdf, err = g.Cdf(1000)
	catchError(t, err)
	expect = 1
	if round(cdf, 6) != expect {
		t.Errorf("A value of Cdf(1000) do not match.\n got: %v\nwant: %v", cdf, expect)
	}

	r := g.Uint64()
	if r < 0 && 1000 < r {
		t.Errorf("An inappropriate value is returned from Rand(). got: %v", r)
	}
}

func TestBenchmarkGamma(t *testing.T) {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := 1000
	k := 0.475
	thita := 170.607

	for n := 2; n < 100000; n *= n {
		_, err := NewGamma(rs, n, k, thita)
		if err != nil {
			t.Logf("%v, n=%v, k=%v, thita=%v", err, n, k, thita)
			break
		}
	}

	thita = 1
	for k := 1; k < 1000; k++ {
		_, err := NewGamma(rs, n, float64(k), thita)
		if err != nil {
			t.Logf("%v, n=%v, k=%v, thita=%v", err, n, k, thita)
			break
		}

	}

	for thita := 1; thita < 1000; thita++ {
		_, err := NewGamma(rs, n, k, float64(thita))
		if err != nil {
			t.Logf("%v, n=%v, k=%v, thita=%v", err, n, k, thita)
			break
		}

	}
}
