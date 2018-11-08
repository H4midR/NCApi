package controllers

import (
	"NCApi/datamodels"

	"github.com/kataras/iris"
)

// NcController : controller for managing nc command
type NcController struct {
	key string
}

//Get : user comment a post
func (c *NcController) Get(ctx iris.Context) string {
	//p1 := datamodels.Point{X: 1, Y: 1, Z: 0}
	//p2 := p1.Add(&p1)
	//v1 := datamodels.Vector{X: 1, Y: 1, Z: 0}
	//ctx.Write(v1.JSON())
	//v1.Unic()
	//v2 := datamodels.Vector{X: 0, Y: 1, Z: 0}
	//_, v3 := v1.AngBetween(&v2)
	//ctx.Write(v3.JSON())
	//redres := math.Acos(v1.Dot(&v2) / v2.Length())
	e1 := datamodels.Polynomial{Coefficient: []float64{0, 1, 1}}
	e2 := datamodels.Polynomial{Coefficient: []float64{1, -1, 2, 1}}
	e3 := e1.Add(&e2)
	e4 := e3.Diff()
	ctx.Write(e3.JSON())
	ctx.Write(e4.JSON())
	return ""
}
