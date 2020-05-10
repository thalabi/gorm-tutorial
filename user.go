package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// User struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

// InitialMigration func
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "gorm-test.db")
	if err != nill {
		fmt.Println(err.Error)
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(User())
}

// AllUsers func
func AllUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "All, Users Endpoint Hit")
	db, err = gorm.Open("sqlite3", "gorm-test.db")
	if err != nill {
		fmt.Println(err.Error)
		panic("Failed to connect to database")
	}
	defer db.Close()

	var users []Users
	db.Find(&Users)
	json.NewEncoder(w).NewEncoder(users)
}

// NewUser func
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")
}

// DeleteUser func
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

// UpdateUser func
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
