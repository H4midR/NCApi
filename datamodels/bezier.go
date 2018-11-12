package datamodels

import (
	"encoding/json"
	"errors"

	"github.com/kataras/iris"
)

// CPoints : cpoints
type CPoints struct {
	Curve      string  `json:"curve,omitempty"`
	Req        string  `json:"req,omitempty"`
	Points     []Point `json:"points,omitempty"`
	Resolotion int     `json:"resolotion,omitempty"`
}

// Bezier : bezier struct
type Bezier struct {
	CP []Point               `json:"cp"`
	BF []BernsteinPolynomial `json:"basicFunctions"`
	U  float64               `json:"u"`
	N  uint64                `json:"n"`
}

//BernsteinPolynomial : bezier basic functions
type BernsteinPolynomial struct {
	Pu  Polynomial `json:"p(u)"`
	Pui Polynomial `json:"p(u-1)"`
	N   uint64     `json:"n"`
	I   uint64     `json:"i"`
}

///----------------------------------------------------------------helper function

//FactorialMemoization
func Factorial(n uint64) (res uint64) {

	if n > 0 {
		res = n * Factorial(n-1)
		return res
	}

	return 1
}

// Init : initialize the Bernstein Polynomial
func (bp *BernsteinPolynomial) Init() {
	var PuCof = make([]float64, bp.N+1)
	var PuiCof = make([]float64, bp.N+1)
	PuCof[bp.I] = float64(Factorial(bp.N) / (Factorial(bp.I) * Factorial(bp.N-bp.I)))
	PuiCof[bp.N-bp.I] = 1

	bp.Pu = Polynomial{Coefficient: PuCof}
	bp.Pui = Polynomial{Coefficient: PuiCof}
}

// Cal : calculate the value of Bu,i(u)
func (bp *BernsteinPolynomial) Cal(u float64) float64 {
	return bp.Pu.Cal(u) * bp.Pui.Cal(1-u)
}

//Init : Init bezier
func (b *Bezier) Init(ctx iris.Context) error {
	b.N = uint64(len(b.CP) - 1)
	if b.N <= 0 {
		ctx.Writef("{\"Message\":\"control points must be more than 1\",\"status\":\"error\"}")
		return errors.New("control points must be more than 1")
	}
	b.BF = make([]BernsteinPolynomial, b.N+1)
	var i uint64
	for i = 0; i <= b.N; i++ {
		b.BF[i] = BernsteinPolynomial{N: b.N, I: i}
		b.BF[i].Init()
	}
	return nil
}

//JSON : string json
func (b *BernsteinPolynomial) JSON() []byte {
	res, _ := json.Marshal(b)
	return res
}

//Cal : calculate the curve point at u
func (b *Bezier) Cal(u float64) Point {
	var i uint64
	res := Point{}
	for i = 0; i <= b.N; i++ {
		p := b.CP[i].Multiplication(b.BF[i].Cal(u))
		res.SAdd(&p)
	}
	return res
}
