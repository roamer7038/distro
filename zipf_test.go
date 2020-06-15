package distro

import (
	"math/rand"
	"testing"
	"time"
)

func TestZipf(t *testing.T) {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	z, err := NewZipf(rs, 1000, 0.693)
	catchError(t, err)

	pdf, err := z.Pdf(1)
	catchError(t, err)

	cdf, err := z.Cdf(1)
	catchError(t, err)

	t.Log(pdf, cdf)
	if pdf != cdf {
		t.Errorf("A value of Pdf(1) and Cdf(1) do not match.\n pdf: %v\ncdf: %v", pdf, cdf)
	}

	expect := 0.040888100707248
	if round(pdf, 15) != expect {
		t.Errorf("A value of Pdf(1) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	pdf, err = z.Pdf(1000)
	catchError(t, err)
	t.Log(cdf)
	expect = 0.000340876402345
	if round(pdf, 15) != expect {
		t.Errorf("A value of Pdf(1000) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	cdf, err = z.Cdf(1000)
	catchError(t, err)
	t.Log(cdf)
	expect = 1
	if round(cdf, 12) != expect {
		t.Errorf("A value of Cdf(1000) do not match.\n got: %v\nwant: %v", cdf, expect)
	}

	r := z.Uint64()
	t.Log(r)
	if r < 0 && 1000 < r {
		t.Errorf("An inappropriate value is returned from Rand(). got: %v", r)
	}
}
