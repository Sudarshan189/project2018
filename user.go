package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"encoding/json"
)





type Users struct {
	ID 		  uint 		`gorm:"primary_key" json:"id"` 		// Aadhar ID
	Name      string 	`gorm:"type:varchar(20)" json:"name"`	// Name
	Address   string 	`gorm:"type:varchar(30)" json:"address"`	// Address
	Phone 	  uint 		`gorm:"not null" json:"phone"`			// Phone No
	Email 	  string 	`gorm:"type:varchar(20); not null" json:"email"`	// Eemail
	RRNum     string 	`gorm:"type:varchar(10)" json:"rr_num"`	// RR Number
	MeterNo   uint 		`gorm:"unique;not null" json:"meter_no"`	// Meter Num
	Tariff 	  string	`json:"tariff"`							// Tarif plan
	SanctLoad string	`json:"sanct_load"`						// Sanctioned Laod
	Password  string	`gorm:"not null" json:"password"`			// Password
}

type Login struct {
	ID uint `json:"id"`
	Password string `json:"password"`
}




var login Login

func Createuser(w http.ResponseWriter, r *http.Request) {
	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	db, err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err != nil {
		json.NewEncoder(w).Encode(DatabaseError)
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Users{}) // migrates the table
	db.Create(&user)
	Showlog(w,r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreatedSuccess)
}


func Getuser(w http.ResponseWriter, r *http.Request) {
	var user Users
	_=json.NewDecoder(r.Body).Decode(&login)
	db,err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err != nil {
		json.NewEncoder(w).Encode(DatabaseError)
		panic(err.Error())
	}
	defer db.Close()
	Showlog(w,r)
	db.Where(&Users{ID:login.ID, Password:login.Password}).First(&user)
	w.Header().Set("Content-Type", "application/json")
	if user.ID != 0 {
		json.NewEncoder(w).Encode(LoggedInSuccess)
	} else {
		json.NewEncoder(w).Encode(LoggedInFail)
	}
}

func Updateuser(w http.ResponseWriter, r *http.Request) {

}
