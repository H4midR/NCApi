package datamodels

// Bezier : bezier struct
type Bezier struct {
	CP []Point `json:"cp"`
	U  float64 `json:"u"`
}

type BernsteinPolynomial struct {
	Pu  Polynomial `json:"p(u)"`
	Pui Polynomial `json:"p(u-1)"`
	N   uint64     `json:"n"`
	I   uint64     `json:"i"`
}

///helper function
//FactorialMemoization
func Factorial(n uint64) (res uint64) {

	if n > 0 {
		res = n * Factorial(n-1)
		return res
	}

	return 1
}

func (bp *BernsteinPolynomial) init() {
	var PuCof = make([]float64, bp.N)
	var PuiCof = make([]float64, bp.N)
	PuCof[bp.I] = float64(Factorial(bp.N) / (Factorial(bp.I) * Factorial(bp.N-bp.I)))
	PuiCof[bp.N-bp.I] = 1

	bp.Pu = Polynomial{Coefficient: PuCof}
	bp.Pui = Polynomial{Coefficient: PuiCof}
}
