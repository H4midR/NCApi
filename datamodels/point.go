package datamodels

import (
	"encoding/json"
	"math"
)

// Point : point ; a 3D Object
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// JSON : convert a point to string json format						Done
func (p *Point) JSON() []byte {
	res, _ := json.Marshal(p)
	return res
}

// Add : add to point and return result								Done
func (p *Point) Add(p2 *Point) Point {
	res := Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}

	return res
}

//SAdd : add a point to itself
func (p *Point) SAdd(p2 *Point) {
	p.X = p.X + p2.X
	p.Y = p.Y + p2.Y
	p.Z = p.Z + p2.Z
}

// Move : move a point to new cordinate								Done
func (p *Point) Move(X float64, Y float64, Z float64) {
	p.X = X
	p.Y = Y
	p.Z = Z
}

// Distance : calculate distance between to Point					Done
func (p *Point) Distance(p2 *Point) float64 {
	dx := math.Pow(p.X-p2.X, 2)
	dy := math.Pow(p.Y-p2.Y, 2)
	dz := math.Pow(p.Z-p2.Z, 2)
	res := math.Sqrt(dx + dy + dz)
	return res
}

// Vector : convert Point to vector									Done
func (p *Point) Vector() Vector {
	res := Vector{X: p.X, Y: p.Y, Z: p.Z}
	return res
}

//Clone : copy the vector										Done
func (p *Point) Clone() Point {
	vh := Point{X: p.X, Y: p.Y, Z: p.Z}
	return vh
}

// SMultiplication : Multiplication with a number itself
func (p *Point) Multiplication(val float64) Point {
	return Point{
		X: p.X * val,
		Y: p.Y * val,
		Z: p.Z * val,
	}
}

// SMultiplication : Multiplication with a number itself
func (p *Point) SMultiplication(val float64) {
	p.X = p.X * val
	p.Y = p.Y * val
	p.Z = p.Z * val
}
