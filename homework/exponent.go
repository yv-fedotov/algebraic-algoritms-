package homework

import "math"

type Exponent struct {
}

func NewExponent() *Exponent {
	return &Exponent{}
}

func (e *Exponent) Exp(a float64, n int) float64 {
	result := 1.0
	for i := 0; i < n; i++ {
		result = result * a
	}

	return math.Floor(result*100000000000) / 100000000000
}

func (e *Exponent) Mul(a float64, n int) float64 {
	switch {
	case n == 0:
		return 1.0
	case n%2 == 0:
		return e.Mul(a, n/2) * e.Mul(a, n/2)
	case n%2 == 1:
		return e.Mul(a, n-1) * a
	default:
		return 0.0
	}
}
