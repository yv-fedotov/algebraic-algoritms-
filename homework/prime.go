package homework

type Prime struct {
}

func NewPrime() *Prime {
	return &Prime{}
}

func (p *Prime) PrimeIter(n int) int {
	prm := make([]int, 0)

	isPrime := func(n int) bool {
		for j := 2; 2*j <= n; j++ {
			if n%j == 0 {
				return false
			}
		}
		return true
	}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			prm = append(prm, i)
		}
	}

	return len(prm)
}

func (p *Prime) PrimeOptima(n int) int {
	prm := make([]bool, n+1)
	if n == 2 {
		return 1
	}
	for i := 2; i <= n; i++ {
		if !prm[i] {
			for k := i * i; k <= n; k = k + i {
				prm[k] = true
			}
		}
	}

	count := 0
	for i := 0; i < len(prm); i++ {
		if !prm[i] {
			count++
		}
	}
	return count - 2
}
