package datamodels

import (
	"encoding/json"
	"errors"
	"log"
	"math"

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
	CP     []Point               `json:"cp"`
	QP     []Point               `json:"qp,omitempty"`
	BF     []BernsteinPolynomial `json:"basicFunctions"`
	BDifF  []BernsteinPolynomial `json:"basicDiffFunctions"`
	U      float64               `json:"u"`
	N      uint64                `json:"n"`
	Length float64               `json:"length"`
}

//BernsteinPolynomial : bezier basic functions
type BernsteinPolynomial struct {
	Pu  Polynomial `json:"p(u)"`
	Pui Polynomial `json:"p(u-1)"`
	N   uint64     `json:"n"`
	I   uint64     `json:"i"`
}

///----------------------------------------------------------------helper function
// JSON : return a string json format of polynominal
func (b *Bezier) JSON() []byte {
	res, _ := json.Marshal(b)
	return res
}

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
	b.QP = make([]Point, b.N)
	b.BDifF = make([]BernsteinPolynomial, b.N)
	var i uint64
	for i = 0; i <= b.N; i++ {
		b.BF[i] = BernsteinPolynomial{N: b.N, I: i}
		b.BF[i].Init()

		//		calculate Q[i]
		if i < b.N {
			tmp1Point := b.CP[i].Minus()
			tmp2Point := b.CP[i+1].Add(&tmp1Point)
			tmp2Point.SMultiplication(float64(b.N))
			b.QP[i] = tmp2Point.Clone()
			b.BDifF[i] = BernsteinPolynomial{N: b.N - 1, I: i}
			b.BDifF[i].Init()
		}

	}

	l := b.LengthCal(10000) // what is the best resulution for calculation of curve length? maybe it's better to be dependent on it's value or curveture 1:1000000 , 2:1000
	//log.Printf("B length is %f", l)
	u := 0.1
	log.Printf("ds/du %f @ %f & lenght is %f", b.DsPDu(u, 1000), u, l)

	return nil
}

//JSON : string json
func (bp *BernsteinPolynomial) JSON() []byte {
	res, _ := json.Marshal(bp)
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

//DiffCal : calculate the curve tangent vector at u
func (b *Bezier) DiffCal(u float64) Vector {
	var i uint64
	res := Point{}
	for i = 0; i < b.N; i++ {
		p := b.QP[i].Multiplication(b.BDifF[i].Cal(u))
		res.SAdd(&p)
	}
	return res.Vector()
}

//Length : calculate the length
func (b *Bezier) LengthCal(n uint32) float64 {
	var res float64
	var i uint32
	for i = 0; i <= n; i++ {
		p2 := b.Cal(float64(i+1) / float64(n+1))
		p1 := b.Cal(float64(i) / float64(n+1))
		res += math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2) + math.Pow(p2.Z-p1.Z, 2))
	}
	b.Length = res
	return res
}

//DsPDu : calculation of Ds/Du , @ u float64 with the calculation Resolution of n uint32
func (b *Bezier) DsPDu(u float64, n uint32) float64 {
	p1 := b.Cal(u - 1/float64(2*n))
	p2 := b.Cal(u + 1/float64(2*n))
	return math.Sqrt(math.Pow(p2.X-p1.X, 2)+math.Pow(p2.Y-p1.Y, 2)+math.Pow(p2.Z-p1.Z, 2)) * float64(n)

}
