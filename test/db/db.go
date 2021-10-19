package db

import "log"

// db act as a dummy package level database.
var db map[string]bool

// init initialize a dummy db with some data
func init() {
	db = make(map[string]bool)
	db["ankuranand@dummy.org"] = true
	db["anand@example.com"] = true
	db["bekzhan@example.com"] = true
}

// UserExists check if the User is registered with the provided email.
func UserExistDb(email string) bool {
	log.Println("call db func")
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}