package service

import (
"testing"
"log"
)

var userExistMock func(email string)bool

type preCheckMock struct{}

func (u *preCheckMock) UserExists(email string)bool {
	return userExistMock(email)
}

func TestCheckUserExist(t *testing.T) {

	user2 := User{
		Name : "Bekzhan",
		Email:"sbekzhan@example.com",
		UserName : "pussy",
	}
	regPreCond = &preCheckMock{}

	//1 case, err != nil - user exist
	//case false, fake behavior
	userExistMock = func(email string)bool {
		return true
	}
	//userExistMock("bekzhan@example.com")

if 	err := RegisterUser(user2, regPreCond); err != nil {
		// log.Println(err, "Expected Register User to throw and error got nil")
		t.Fatal(err,1)
	}
//fake behaviror user have
	userExistMock = func(email string)bool {
		return false
	}

if 	err := RegisterUser(user2, regPreCond); err == nil {
		log.Println(err, "Expected Register User to throw and error got nil")
	}
}

// mock - requirement interface
