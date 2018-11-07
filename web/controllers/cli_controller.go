package controllers

import (
	"ehome/datamodels"
	"ehome/datastartpoint"
	"ehome/services"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

/*eslint-disable */
//							work flow : 10%

//CliController : /cli/* - controler for administorators
type CliController struct{}

//Get : null
func (c *CliController) Get() string {
	return "cli"
}

//Get : null
func (c *CliController) GetTest() string {
	return "clitest"
}

//GetCategories : return root categories				DONE
func (c *CliController) GetCategories() string {
	myg := services.NewMygraphService()
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")
	q := fmt.Sprintf(`
		{
			data(func: uid(0x%x)) {
				child @facets(orderdesc: ordinal){
					uid
					expand(_all_)
				}
			}
		}
		`, datastartpoint.CatRoot)
	return string(myg.Query(q))
}

//GetCategoriesBy : return cat node with uid given		DONE
func (c *CliController) GetCategoriesBy(uid string) string {
	myg := services.NewMygraphService()
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")
	UID, _ := strconv.ParseUint(uid, 16, 64)
	q := fmt.Sprintf(`
		{
			data(func: uid(%#x))  @filter(eq(key,"cat")) {
				name
				title
				comment
				href
				ordinal
				child @facets(orderdesc: ordinal){
					uid
					expand(_all_)
				}
			}
		}
		`, UID)
	return string(myg.Query(q))
}

//GetProps : return all avilable porps					DONE
func (c *CliController) GetProps() string {
	myg := services.NewMygraphService()
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")
	q := fmt.Sprintf(`
		{
			data(func: uid(%#x)) {
				uid
				child @facets(orderdesc: ordinal){
					uid
					expand(_all_)
				}
			}
		}
		`, datastartpoint.PropRoot)
	return string(myg.Query(q))
}

//GetPropsName : return all avilable porps uid					DONE
func (c *CliController) GetPropsName(ctx iris.Context) string {
	myg := services.NewMygraphService()
	type Response struct {
		UID   string     `json:"uid"`
		Child []Response `json:"child,omitempty"`
		Name  string     `json:"name,omitempty"`
	}
	var res struct {
		Data []Response `json:"data"`
	}
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")
	q := fmt.Sprintf(`
	 	{
	 		data(func: uid(%#x)) {
	 			uid
	 			child{
	 				uid
	 				name
	 			}
	 		}
	 	}
	 	`, datastartpoint.PropRoot)
	req := myg.Query(q)
	json.Unmarshal(req, &res)
	//str, _ := json.Marshal(res.Data[0].Child)
	for _, element := range res.Data[0].Child {
		//index is the index where we are
		//element is the element from someSlice for where we are
		ctx.Writef("<%s> <name> \"%s\" .\n", element.UID, element.Name)
	}
	return ""
}

//GetByProps : return all avilable porps of node		DONE
func (c *CliController) GetByProps(uid string) string {

	myg := services.NewMygraphService()
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")
	UID, _ := strconv.ParseUint(uid, 16, 64)

	q := fmt.Sprintf(`
		{
			data(func: uid(%#x)){
				uid
				prop{
					uid
					expand(_all_){
						uid
						expand(_all_)
					}
				}
			}
		}
		`, UID)
	return string(myg.Query(q))
}

//GetPropsBy : get a prop ditails
func (c *CliController) GetPropsBy(uid string) string {
	myg := services.NewMygraphService()
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")
	UID, _ := strconv.ParseUint(uid, 16, 64)
	q := fmt.Sprintf(`
		{
			data(func: uid(%#x)) @filter(eq(key,"prop")){
				uid
				expand(_all_){
					uid
					expand(_all_)
				}
			}
		}
		`, UID)
	return string(myg.Query(q))
}

//GetByPropsNot : return all avilable props of node expect porps node has
func (c *CliController) GetByPropsNot(suid string) string {
	myg := services.NewMygraphService()
	//var cat datamodels.Cat
	//ctx.Header("Access-Control-Allow-Origin", "*")
	uid, _ := strconv.ParseUint(suid, 16, 64)
	q := fmt.Sprintf(`
		{
		var(func:uid(%#x)){
			prop{
			reserved as uid
			}
		}
		data(func:uid(%#x)){
			uid
			child @filter(NOT(uid(reserved))){
			uid
			expand(_all_)
			}
		}
		}
		`, uid, datastartpoint.PropRoot)
	return string(myg.Query(q))
}

//PostProps : add new prop
func (c *CliController) PostProps(ctx iris.Context) string {
	//ctx.Header("Access-Control-Allow-Origin", "*")
	var myprop datamodels.Prop
	ctx.ReadForm(&myprop)
	res, _ := json.Marshal(myprop)
	return string(res)
}

//PostByPropBy : add a prop to a node
func (c *CliController) PostByPropBy(ucatid string, upropid string) string {
	cid, _ := strconv.ParseUint(ucatid, 16, 64)
	pid, _ := strconv.ParseUint(upropid, 16, 64)
	str := fmt.Sprintf(`<%#x> <prop> <%#x> . `, cid, pid)
	myg := services.NewMygraphService()
	q := []byte(str)
	type result struct {
		Message string `json:"message,omitempty"`
		State   string `json:"state,omitempty"`
	}
	res := result{
		Message: string(myg.MutateRDF(q, "add")),
		State:   "ok",
	}
	r, _ := json.Marshal(res)
	return string(r)
}

//PostCatByPropByDelete : delete a prop from a cat
func (c *CliController) PostByPropByDelete(ucatid string, upropid string) string {
	cid, _ := strconv.ParseUint(ucatid, 16, 64)
	pid, _ := strconv.ParseUint(upropid, 16, 64)
	str := fmt.Sprintf(`<%#x> <prop> <%#x> . `, cid, pid)
	myg := services.NewMygraphService()
	q := []byte(str)
	type result struct {
		Message string `json:"message,omitempty"`
		State   string `json:"state,omitempty"`
	}
	res := result{
		Message: string(myg.MutateRDF(q, "delete")),
		State:   "ok",
	}
	r, _ := json.Marshal(res)
	return string(r)
}

//PostObjByPropBy : add a prop to a obj
func (c *CliController) PostObjByPropBy(uobjid string, upropid string) string {
	q := `Objid : %x \n propid : %x`
	oid, _ := strconv.ParseUint(uobjid, 16, 64)
	pid, _ := strconv.ParseUint(upropid, 16, 64)
	str := fmt.Sprintf(q, oid, pid)
	return string(str)
}

//GetNameBy : tmp method
func (c *CliController) GetNameBy(name string) string {
	myg := services.NewMygraphService()
	q := `query all($name: string) {
    all(func: eq(name, $name)) {
		uid
		name
    }
  }`
	return string(myg.VQuery(q, map[string]string{"$name": "hamid"}))

}
