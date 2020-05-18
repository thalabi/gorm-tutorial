package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// User struct
type User struct {
	gorm.Model
	Name         string
	Email        string
	PhoneNumbers []PhoneNumber
}

// PhoneNumber struct
type PhoneNumber struct {
	gorm.Model
	Phone  string
	Number string
	UserID uint
}

// InitialMigration func
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "gorm-test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&PhoneNumber{})
}

// AllUsers func
func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "gorm-test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	var users []User
	db.Debug().Preload("PhoneNumbers").Find(&users)
	// Retrieve phone numbers for user
	// for i, user := range users {
	// 	var phoneNumbers []PhoneNumber
	// 	db.Where("user_id", user.ID).Find(&phoneNumbers)
	// 	users[i].PhoneNumbers = phoneNumbers
	// }
	fmt.Println("users: ", users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// AddUser func
func AddUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "gorm-test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
	cellPhone := PhoneNumber{Phone: "Cell", Number: r.Form["Cell"][0]}
	homePhone := PhoneNumber{Phone: "Home", Number: r.Form["Home"][0]}
	fmt.Println("homePhone ", homePhone)
	fmt.Println("cellPhone ", cellPhone)

	user := User{Name: name, Email: email, PhoneNumbers: []PhoneNumber{cellPhone, homePhone}}
	//db.Create(&User{Name: name, Email: email, PhoneNumbers: []PhoneNumber{cellPhone, homePhone}})
	db.Create(&user)
	fmt.Println("user: ", user)
	fmt.Fprint(w, "New User Created with name ", name, " and email ", email)
}

// DeleteUser func
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "gorm-test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Debug().Delete(&user)

	fmt.Fprint(w, "User Successfully Deleted, name ", name)
}

// UpdateUser func
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "gorm-test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	fmt.Fprint(w, "Updated User name ", name, " email ", email)
}
