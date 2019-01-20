package controllers

import (
	"ehome/datamodels"
	"ehome/datastartpoint"
	"ehome/services"
	"encoding/json"
	"log"
)

/*eslint-disable */

//AdminController : /admin/* - controler for administorators
type AdminController struct{}

//Get : null
func (c *AdminController) Get() string {
	myg := services.NewMygraphService()
	q := []byte(`<0x1f> <close> <0x2f> .
	`)
	return string(myg.MutateRDF(q, "add"))

	return "admin"
}

//GetMutate1 : tmp action method
func (c *AdminController) GetMutate1() string {
	myg := services.NewMygraphService()
	type RootProp struct {
		UID   uint64            `json:"uid,omitempty"`
		Name  string            `json:"name,omitempty"`
		Child []datamodels.Prop `json:"child,omitempty"`
	}
	prop := RootProp{
		UID:   datastartpoint.PropRoot,
		Child: []datamodels.Prop{
			// {
			// 	Name:         "نوع سازه",
			// 	Comment:      "نوع سازه",
			// 	Quantum:      "",
			// 	Type:         "string",
			// 	Min:          0,
			// 	Max:          1000,
			// 	TagInterface: "mytext",
			// 	Key:          "prop",
			// 	Kind:         "progenitor",
			// 	Icon:         "mdi-chart-bubble",
			// },
			// {
			// 	Name:         "نوع سقف",
			// 	Comment:      "نوع سقف",
			// 	Quantum:      "",
			// 	Type:         "string",
			// 	Min:          0,
			// 	Max:          1000,
			// 	TagInterface: "mytext",
			// 	Key:          "prop",
			// 	Kind:         "progenitor",
			// 	Icon:         "mdi-chart-bubble",
			// },
			// {
			// 	Name:         "نوع ملک",
			// 	Comment:      "نوع ملک",
			// 	Quantum:      "",
			// 	Type:         "extra",
			// 	TagInterface: "myextra",
			// 	Key:          "prop",
			// 	Kind:         "progenitor",
			// 	Icon:         "mdi-chart-bubble",
			// 	Defaults: []datamodels.PropDefault{
			// 		{
			// 			Title:   "",
			// 			Value:   "",
			// 			Key:     "prop",
			// 			Kind:    "defaults",
			// 			Ordinal: ,
			// 		},

			// 	},
			// },
		},
	}

	q, err := json.Marshal(prop)
	if err != nil {
		log.Fatal(err)
	}
	str, _ := myg.Mutate(q)
	return str

}

//GetNameBy : tmp method
func (c *AdminController) GetNameBy(name string) string {
	myg := services.NewMygraphService()
	q := `query all($name: string) {
    all(func: eq(name, $name)) {
		uid
		name
    }
  }`
	return string(myg.VQuery(q, map[string]string{"$name": "hamid"}))

}
