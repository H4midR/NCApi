package datamodels

import (
	"golang.org/x/crypto/bcrypt"
)

//UserRoot : the root of user
type UserRoot struct {
	UID   uint64 `json:"uid,omitempty"`
	Name  string `json:"name,omitempty"`
	Users []User `json:"child,omitempty"`
}

//User : data model for user node . user node is subnode of one of the three main node . admin . verifier . person. root is temp for unverified user
type User struct {
	UID        uint64 `json:"uid,omitempty" form:"uid"`
	Name       string `json:"name,omitempty" form:"name"`             //name f user
	Mobile     string `json:"mobile,omitempty" form:"mobile"`         //mobile . is primery and uniq
	Email      string `json:"email,omitempty" form:"email"`           //email . not very important
	Password   string `json:"password,omitempty" form:"password"`     //hashed password
	Coname     string `json:"coname,omitempty" form:"coname"`         //if filled user is corporation and the name of corporation is so
	Phone      string `json:"phone,omitempty" form:"phone"`           //static phone
	Bio        string `json:"bio,omitempty" form:"bio"`               //a brief introduction of companys
	Img        string `json:"img,omitempty" form:"img"`               //img or logo of a corporation
	SignupDate string `json:"signupDate,omitempty" form:"signupDate"` //created date
	Verified   bool   `json:"verified,omitempty" form:"verified"`     //is verified user
	MVCode     string `json:"mvCode,omitempty" form:"mvcode"`         //code sent to mobile for verification - Mobile Verifi Code
	Master     uint64 `json:"master,omitempty" form:"master"`         //master user .
	Token      string `json:"token,omitempty" form:"token"`           //token made for every session , session less and jwt
	Key        string `json:"key,omitempty"`
}

// GeneratePassword : will generate a hashed password for us based on the
// user's input.
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword : will check if passwords are matched.
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}
