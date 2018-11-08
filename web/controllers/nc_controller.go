package controllers

import (
	"encoding/json"

	"github.com/kataras/iris"
)

// NcController : controller for managing nc command
type NcController struct {
	key string
}

//Get : user comment a post
func (c *NcController) Get(ctx iris.Context) {
	//p1 := []datamodels.Point{datamodels.Point{X: 0, Y: 0}, datamodels.Point{X: 1, Y: 1}, datamodels.Point{X: 2, Y: 0}}

	// p := datamodels.Point{}
	// p.SAdd(&p1[0])
	// p.SAdd(&p1[1])
	// p.SAdd(&p1[2])
	//b := datamodels.Bezier{CP: p1}
	//b.Init()
	//bp := datamodels.BernsteinPolynomial{N: 3, I: 0}
	//bp.Init()

	//p2 := p1.Add(&p1)
	//v1 := datamodels.Vector{X: 1, Y: 1, Z: 0}
	//ctx.Write(v1.JSON())
	//v1.Unic()
	//v2 := datamodels.Vector{X: 0, Y: 1, Z: 0}
	//_, v3 := v1.AngBetween(&v2)
	//ctx.Write(v3.JSON())
	//redres := math.Acos(v1.Dot(&v2) / v2.Length())
	//e1 := datamodels.Polynomial{Coefficient: []float64{0, 1, 1}}
	//e2 := datamodels.Polynomial{Coefficient: []float64{1, -1, 2, 1}}
	//e3 := e1.Add(&e2)
	//e4 := e3.Diff()
	//ctx.Write(e3.JSON())
	//ctx.Write(e4.JSON())
	//res := b.Cal(0.5)
	//ctx.Write(res.JSON())
	//ctx.View("index.html")
}
func (c *NcController) Post(ctx iris.Context) string {

	type request struct {
		CP []float64 `json:"cp"`
	}
	type Requeststr struct {
		Name string
		Last string
	}
	var req Requeststr
	ctx.ReadJSON(&req)
	js, _ := json.Marshal(req)
	ctx.Write(js)
	ctx.Writef("Received: %#+v\n", req)
	return ""

}
