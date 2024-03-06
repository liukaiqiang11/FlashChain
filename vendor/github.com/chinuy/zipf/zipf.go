package zipf

import (
	"math"
	"math/rand"
)

func NewZipf(r *rand.Rand, s float64, imax uint64) *Zipf {
	z := &Zipf{}
	z.rng = r

	// convert to int for convenience, not sure any side effect
	var n int = int(imax)

	tmp := make([]float64, n)
	for i := 1; i < n+1; i++ {
		tmp[i-1] = 1.0 / math.Pow(float64(i), s)
	}

	zeta := make([]float64, n+1)
	distMap := make([]float64, n+1)
	z.prob = make([]float64, n)

	zeta[0] = 0
	for i := 1; i < n+1; i++ {
		zeta[i] += zeta[i-1] + tmp[i-1]
	}

	for i := 0; i < n+1; i++ {
		distMap[i] = zeta[i] / zeta[n]
	}

	for i := 1; i < n+1; i++ {
		z.prob[i-1] = distMap[i] - distMap[i-1]
	}

	return z
}

type Zipf struct {
	s float64
	prob []float64
	rng *rand.Rand
}

func (z *Zipf) Uint64() uint64 {
	u := z.rng.Float64()
	for i := 0; i < len(z.prob); i++ {
		if u < z.prob[i] {
			return uint64(i)
		}
		u -= z.prob[i]
	}
	// should not come here
	// return last index for safty
	return uint64(len(z.prob)-1)
}
