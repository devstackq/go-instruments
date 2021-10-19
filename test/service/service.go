package service

import (
	"fmt"
	"testexample/db"
)

type RegPreCheck struct{}

var regPreCond registrationPreChecker

func init(){
	regPreCond = &RegPreCheck{}
}

// User encapsulate a user in the system.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

type registrationPreChecker interface {
	UserExists(string)bool
}

func (r *RegPreCheck)UserExists(email string)bool {
	return db.UserExistDb(email)
}

func NewRegistrationPreChecker () registrationPreChecker { 
	return &RegPreCheck {} 
 } 

// RegisterUser will register a User if only User has not been previously
// registered.
func RegisterUser(user User, r registrationPreChecker) error {
// check if user is already registered
//	found := db.UserExists(user.Email)
found := r.UserExists(user.Email)

if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	// carry business logic and Register the user in the system
	// log.Println(user)
	return nil
}