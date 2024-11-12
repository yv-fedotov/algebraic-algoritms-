package homework

import "math"

type Fib struct {
	exp *Exponent
}

func NewFib() *Fib {
	return &Fib{
		exp: NewExponent(),
	}
}

func (f *Fib) FibonacciRec(n int) uint64 {
	if n < 2 {
		return 1
	}

	return f.FibonacciRec(n-1) + f.FibonacciRec(n-2)
}

func (f *Fib) FibonacciIter(n int) uint64 {
	sum, j := uint64(0), uint64(1)
	for i := 0; i < n; i++ {
		sum += j
		j = sum - j
	}

	return sum
}

func (f *Fib) FibonacciGold(n int) uint64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1

	default:
		sqr5 := math.Sqrt(5)
		fi := (1.0 + sqr5) / 2.0
		result := f.exp.Mul(fi, n)/sqr5 + 1.0/2.0
		return uint64(result)
	}
}
