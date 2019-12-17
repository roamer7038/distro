package distro

import (
	"github.com/roamer7038/distro/internal/mathutils"
	"math/rand"
	"testing"
	"time"
)

func TestZipf(t *testing.T) {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	z, _ := NewZipf(rs, 1000, 0.693)

	pdf, _ := z.Pdf(1)
	cdf, _ := z.Cdf(1)
	if pdf != cdf {
		t.Errorf("A value of Pdf(1) and Cdf(1) do not match.\n pdf: %v\ncdf: %v", pdf, cdf)
	}

	expect := 0.040888100707248
	if mathutils.Round(pdf, 15) != expect {
		t.Errorf("A value of Pdf(1) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	pdf, _ = z.Pdf(1000)
	expect = 0.000340876402345
	if mathutils.Round(pdf, 15) != expect {
		t.Errorf("A value of Pdf(1000) do not match.\n got: %v\nwant: %v", pdf, expect)
	}

	cdf, _ = z.Cdf(1000)
	expect = 1
	if mathutils.Round(cdf, 12) != expect {
		t.Errorf("A value of Cdf(1000) do not match.\n got: %v\nwant: %v", cdf, expect)
	}

	r := z.Uint64()
	if r < 0 && 1000 < r {
		t.Errorf("An inappropriate value is returned from Rand(). got: %v", r)
	}
}
