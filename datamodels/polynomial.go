package datamodels

import (
	"encoding/json"
	"math"
)

// Polynomial : a polynomial Equation
type Polynomial struct {
	Coefficient []float64 `json:"coefficient"`
}

// Add : add to polynomial function
func (e *Polynomial) Add(e2 *Polynomial) Polynomial {
	l1 := float64(len(e.Coefficient))
	l2 := float64(len(e2.Coefficient))
	l := int(math.Max(l1, l2))
	Coef := make([]float64, l)
	li1 := int(l1)
	li2 := int(l2)
	for i := 0; i < l; i++ {
		Coef[i] = 0
		if i < li1 {
			Coef[i] += e.Coefficient[i]
		}
		if i < li2 {
			Coef[i] += e2.Coefficient[i]
		}
	}
	res := Polynomial{Coefficient: Coef}
	return res
}

// JSON : return a string json format of polynominal
func (e *Polynomial) JSON() []byte {
	res, _ := json.Marshal(e)
	return res
}

// Cal : calculate the value of function at u
func (e *Polynomial) Cal(u float64) float64 {
	l := len(e.Coefficient)
	var res float64
	res = 0
	for i := 0; i < l; i++ {
		res += e.Coefficient[i] * math.Pow(u, float64(i))
	}
	return res
}

// Diff : return diffrensial polynomial
func (e *Polynomial) Diff() Polynomial {
	l := len(e.Coefficient)
	var res = make([]float64, l)
	for i := 0; i < l-1; i++ {
		res[i] = e.Coefficient[i+1] * float64(i+1) // += e.Coefficient[i] * math.Pow(u, float64(i))
	}

	return Polynomial{Coefficient: res}

}
