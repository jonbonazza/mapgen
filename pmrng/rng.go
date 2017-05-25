package pmrng

import (
	"math"
)

type Rand struct {
	seed uint64
}

func (r *Rand) NextUint64() uint64 {
	return r.generate()
}

func (r *Rand) NextFloat64() float64 {
	return float64(r.generate()) / 2147483647.0
}

func (r *Rand) NextUint64Ranger(min, max float64) uint64 {
	min -= .4999
	max += .4999
	return r.round(min + ((max - min) * r.NextFloat64()))
}

func (r *Rand) NextFloat64Range(min, max float64) float64 {
	return min + ((max - min) * r.NextFloat64())
}

func (r *Rand) generate() uint64 {
	r.seed = (r.seed * 16807) % 2147483647
	return r.seed
}

func (r *Rand) round(val float64) uint64 {
	return uint64(math.Floor(val + .5))
}

func New(seed uint64) *Rand {
	return &Rand{
		seed: seed,
	}
}
