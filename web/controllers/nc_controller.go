package controllers

import (
	"NCApi/datamodels"
	"encoding/json"

	"github.com/kataras/iris"
)

// NcController : controller for managing nc command
type NcController struct {
	key string
}

//Get : user comment a post
func (c *NcController) Get(ctx iris.Context) {
	p1 := []datamodels.Point{datamodels.Point{X: 0, Y: 0}, datamodels.Point{X: 1, Y: 1}, datamodels.Point{X: 2, Y: 0}}

	// p := datamodels.Point{}
	// p.SAdd(&p1[0])
	// p.SAdd(&p1[1])
	// p.SAdd(&p1[2])
	b := datamodels.Bezier{CP: p1}
	b.Init(ctx)
	//v := b.DiffCal(1)
	//cp := b.Cal(0.5)
	//ctx.Write(v.JSON())
	b.Go()
	//ctx.Writef("<br>")
	//ctx.Write(cp.JSON())
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

	type CPoints struct {
		Curve      string             `json:"curve,omitempty"`
		Req        string             `json:"req,omitempty"`
		Points     []datamodels.Point `json:"points"`
		Resolotion int                `json:"resolotion,omitempty"`
	}
	var cps CPoints
	ctx.ReadJSON(&cps)
	js, _ := json.Marshal(cps)
	ctx.Write(js)
	//ctx.Writef("Received: %s\n", req.JSON)
	return ""

}

//PostPlot : plot the curve
func (c *NcController) PostPlot(ctx iris.Context) {
	var cps datamodels.CPoints
	ctx.ReadJSON(&cps)
	Bezier := datamodels.Bezier{CP: cps.Points}
	err := Bezier.Init(ctx)
	if err != nil {
		return
	}
	var step, u float64
	step = 1
	step = 1 / float64(cps.Resolotion)
	u = 0
	resPoints := make([]datamodels.Point, cps.Resolotion+1)
	index := 0
	for u = 0; u < 1+step; u += step {
		resPoints[index] = Bezier.Cal(u)
		index++
	}
	res, _ := json.Marshal(resPoints)
	ctx.Write(res)
}

//PostGo : pass the curve
func (c *NcController) PostGo(ctx iris.Context) {
	var cps datamodels.CPoints
	ctx.ReadJSON(&cps)
	Bezier := datamodels.Bezier{CP: cps.Points}
	err := Bezier.Init(ctx)
	if err != nil {
		return
	}
	Bezier.Go()

	ctx.WriteString("done")
}

func (c *NcController) PostTest(ctx iris.Context) {

	type Person struct {
		Name  string `json:"Name,omitempty"`
		Age   int    `json:"Age,omitempty"`
		City  string `json:"City,omitempty"`
		Other string `json:"Other,omitempty"`
	}
	var persons Person
	err := ctx.ReadJSON(&persons)

	if err != nil {
		//ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	str, _ := json.Marshal(persons)

	ctx.Write(str)
}
