package controllers

import (
	"ehome/datamodels"
	"ehome/datastartpoint"
	"ehome/services"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

/*eslint-disable */
// token needed in all trancactions

// UserController : /user/* - controler
type UserController struct{}

//###########################################################################
//																			#
//										GET									#
//																			#
//###########################################################################

//Get : return signin or singup page
func (c *UserController) Get() string {
	t := time.Now()
	SignupDate := t.Format(time.RFC3339)

	return string(SignupDate)
}

//GetBy : get all data of user		- 	uid user id
func (c *UserController) GetBy(uid string) string {
	myg := services.NewMygraphService()
	q := fmt.Sprintf(`
		{
			user(func: uid(0x%s)) @filter(eq(key,"user")) {
				expand(_all_)
			}
		}
		`, uid)
	return string(myg.Query(q))
}

//GetByFavorite : get all favorites post 	-	uid user id
func (c *UserController) GetByFavorite(uid string) string {
	myg := services.NewMygraphService()
	q := fmt.Sprintf(`
		{
			user(func: uid(0x%s)) @filter(eq(key,"user")) {
				favorite @facets(orderdesc: creationDate) {
					uid
					expand(_all_){
						expand(_all_)
					}
				}
			}
		}
		`, uid)
	return string(myg.Query(q))
}

//Post : new user
func (c *UserController) Post(ctx iris.Context) string { // must cheack uniq mobile
	//ctx.Header("Access-Control-Allow-Origin", "*")
	var Myuser datamodels.User
	ctx.ReadForm(&Myuser)
	if Myuser.UID != 0 || Myuser.MVCode != "" || Myuser.Token != "" || Myuser.SignupDate != "" || Myuser.Verified != false {
		return "403 , Forbiden path"
	}
	if Myuser.Mobile == "" || Myuser.Password == "" {
		return "invalid request , invalid Mobile , invalid password"
	}
	myg := services.NewMygraphService()
	HPass, err := datamodels.GeneratePassword(Myuser.Password)
	if err != nil {
		log.Fatal(err)
	}
	Myuser.Password = string(HPass)
	Myuser.MVCode = "1234" // const value, have to be random
	Myuser.Verified = false
	Myuser.Key = "user"
	t := time.Now()
	Myuser.SignupDate = t.Format(time.RFC3339)

	URoot := datamodels.UserRoot{
		UID: datastartpoint.UserRoot,
		Users: []datamodels.User{
			Myuser,
		},
	}

	q, _ := json.Marshal(URoot)
	str, _ := myg.Mutate(q)
	return str
}

//PostBy : edite user
func (c *UserController) PostBy(ctx iris.Context, uid string) string {
	var Myuser struct {
		Name   string `json:"name,omitempty" form:"name"`     //name f user
		Email  string `json:"email,omitempty" form:"email"`   //email . not very important
		Coname string `json:"coname,omitempty" form:"coname"` //if filled user is corporation and the name of corporation is so
		Phone  string `json:"phone,omitempty" form:"phone"`   //static phone
		Bio    string `json:"bio,omitempty" form:"bio"`       //a brief introduction of companys
		Img    string `json:"img,omitempty" form:"img"`       //img or logo of a corporation
		MVCode string `json:"mvCode,omitempty" form:"mvcode"` //code sent to mobile for verification - Mobile Verifi Code
		Token  string `json:"token,omitempty" form:"token"`   //token made for every session , session less and jwt
	}
	ctx.ReadForm(&Myuser)
	UID, _ := strconv.ParseUint(uid, 16, 64)

	//
	//to kind of edite
	//	* : after pre regester , with uid and vmcode wich verify user
	//	* : fillig or edit user data , with uid and token

	if Myuser.MVCode != "" { // after pre regester , with uid and vmcode wich verify user
		//vmc := Myuser.MVCode

		myg := services.NewMygraphService()

		q := `{
					me(func:uid(%#x)) @filter(eq(key,"user")){
						mvCode
						verified
						token
					}
				}
				`
		q = fmt.Sprintf(q, UID)

		res := myg.Query(q)

		var decode struct {
			Me []struct {
				UID      uint64 `json:"uid,omitempty"`
				MVCode   string `json:"mvCode,omitempty"`
				Verified bool   `json:"verified,omitempty"`
				Token    string `json:"token,omitempty"`
			}
		}
		err := json.Unmarshal(res, &decode)
		if err != nil {
			log.Fatal(err)
		}

		if decode.Me[0].Verified {
			return "this user verified before"
		} else {
			if decode.Me[0].MVCode == Myuser.MVCode {
				var reqry datamodels.User
				reqry.UID = UID
				reqry.Verified = true
				//token := [256]byte{}
				//rand.Read(token[:])
				//reqry.Token = string(token[:])
				reqry.Token = "1234567890abc"
				out, _ := json.Marshal(reqry)
				ctx.WriteString(string(out[:]))
				return "hi"
				//return myg.Mutate(out) // must return the token

			} else {

				return "wrong number"
			}
		}

	} else if Myuser.Token != "" { // fillig or edit user data , with uid and token
		myg := services.NewMygraphService()

		q := `{
					me(func:uid(%#x) @fillter(eq(key,"user"))){
						mvCode
						verified
						token
					}
				}
				`
		q = fmt.Sprintf(q, uid)
		res := myg.Query(q)
		var decode struct {
			Me []datamodels.User
		}
		json.Unmarshal(res, &decode)

		switch decode.Me[0].Verified {
		case true:
			if Myuser.Token == decode.Me[0].Token {
				// code here
				return " good lets keep moving"
			}
			return " invalid token "

		case false:
			return "please verify your mobile first"
		}
	}

	return "end of method . nothing happend"
}

//PostFavoriteBy : mark a post as favorite for user	-	uid user id
func (c *UserController) PostByFavoriteBy(uid string, pid string) string {
	// var input struct {
	// 	Token string `json:"token"`
	// 	Action string `json:"action"`
	// }
	// ctx.ReadForm(&input)

	return ""
}

//PostLikeBy : user like a post or unlike
func (c *UserController) PostLikeBy(uid uint64) string {
	return ""
}

//PostCommentBy : user comment a post
func (c *UserController) PostCommentBy(uid uint64) string {
	return ""
}

//PostStarBy : user star a user
func (c *UserController) PostStarBy(uid uint64) string {
	return ""
}

//PostByWallet : user wallet trancaction
func (c *UserController) PostByWallet(uid uint64) string {
	return ""
}

//PostByCharge : charge the wallet user
func (c *UserController) PostByCharge(uid uint64) string {
	return ""
}
