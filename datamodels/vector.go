package datamodels

import (
	"encoding/json"
	"math"
)

// Vector : 3D vector
type Vector struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Start Point   `json:"start,omitempty"`
}

// Add : add tow vector											Done
func (v *Vector) Add(v2 *Vector) Vector {
	res := Vector{
		X:     v.X + v2.X,
		Y:     v.Y + v2.Y,
		Z:     v.Z + v2.Z,
		Start: v.Start,
	}
	return res
}

// Minus : minus all parameters
func (v *Vector) Minus() {}

// Unic : make vector Unic										Done
func (v *Vector) Unic() {
	len := v.Length()
	v.X = v.X / len
	v.Y = v.Y / len
	v.Z = v.Z / len
}

// JSON : return string json format								Done
func (v *Vector) JSON() []byte {
	res, _ := json.Marshal(v)
	return res
}

// Dot : dot tow vector											Done
func (v *Vector) Dot(v2 *Vector) float64 {
	return v.X*v2.X + v.Y + v2.Y + v.Z + v2.Z
}

// SMultiplication : Multiplication with a number itself
func (v *Vector) SMultiplication(val float64) {
	v.X = v.X * val
	v.Y = v.Y * val
	v.Z = v.Z * val
}

// Cross : Cross mutiplation									Done
func (v *Vector) Cross(v2 *Vector) Vector {

	h1 := v.Y*v2.Z - v2.Y*v.Z
	h2 := v2.X*v.Z - v.X*v2.Z
	h3 := v.X*v2.Y - v2.X*v.Y

	return Vector{X: h1, Y: h2, Z: h3, Start: v.Start}
}

// AngBetween : Angle between to vector							unCorrect
func (v *Vector) AngBetween(v2 *Vector) (float64, Vector) {
	vh := v.Clone()
	vh.Unic()
	redres := math.Acos(vh.Dot(v2) / v2.Length())
	vh = vh.Cross(v2)
	vh.Unic()
	vh.SMultiplication(redres)
	return redres, vh
}

// Length : Magnitude of vector									Done
func (v *Vector) Length() float64 {
	dx := math.Pow(v.X, 2)
	dy := math.Pow(v.Y, 2)
	dz := math.Pow(v.Z, 2)
	res := math.Sqrt(dx + dy + dz)
	return res
}

//Clone : copy the vector										Done
func (v *Vector) Clone() Vector {
	vh := Vector{X: v.X, Y: v.Y, Z: v.Z}
	return vh
}
