package controllers

/*eslint-disable */

import (
	"ehome/datamodels"
	"encoding/json"

	"github.com/kataras/iris"
)

//PosterController : /poster/* 		- controler
type PosterController struct{}

//Get : get last 20 poster
func (c *PosterController) Get() string {
	return "hi"
}

//GetBy : get all data of some post		-uid poster id
func (c *PosterController) GetBy(uid uint64) string {
	return "hi"
}

//GetVerify : get all unverified post
func (c *PosterController) GetVerify() string {
	return "hi"
}

//GetVerifyBy : get all unverified post linked to uid - uid mostly place node
func (c *PosterController) GetVerifyBy(uid uint64) string {
	return "hi"
}

//Post : new poster
func (c *PosterController) Post(ctx iris.Context) string {
	//myg := services.NewMygraphService()

	req := datamodels.PosterForm{}
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")

	err := ctx.ReadForm(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	//pid, _ := strconv.ParseUint(upropid, 16, 64)
	//str := fmt.Sprintf(`<%#x> <prop> <%#x> . `, cid, pid)

	str, _ := json.Marshal(req)
	ctx.Write(str)

	//ctx.Write(req.Price)
	//str, _ := json.Marshal(res.Data[0].Child)

	return ""
}

//PostBy : edit poster 					- uid poster id
func (c *PosterController) PostBy(uid uint64) string {
	return "hi"
}

//PostVerifyBy : verify poster 					-   uid poster id
func (c *PosterController) PostVerifyBy(uid uint64) string {
	return "hi"
}

//PostDropBy : delete poster
func (c *PosterController) PostDropBy(uid uint64) string {
	return "hi"
}
