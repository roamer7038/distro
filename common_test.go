package distro

import (
	"math"
	"testing"
)

func catchError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
