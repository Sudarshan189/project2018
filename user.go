package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)




type Users struct {
	ID 		  uint		`gorm:"primary_key" json:"id"` 		// Aadhar ID = Account ID
	Name      string 	`gorm:"type:varchar(50)" json:"name"`	// Name
	Address   string 	`gorm:"type:varchar(100)" json:"address"`	// Address
	Phone 	  uint 		`gorm:"not null type:bigint(20)" json:"phone"`			// Phone No
	Email 	  string 	`gorm:"type:varchar(50); not null" json:"email"`	// Eemail
	RRNum     string 	`gorm:"type:varchar(20);unique;not null" json:"rr_num"`	// RR Number
	MeterNo   uint 		`gorm:"unique;not null" json:"meter_no"`	// Meter Num
	Tariff 	  string	`json:"tariff"`							// Tarif plan
	SanctLoad string	`json:"sanct_load"`						// Sanctioned Laod
	Password  string	`gorm:"not null" json:"password"`			// Password
}

type Login struct {
	ID uint `json:"id"`
	Password string `json:"password"`
}

type ChangePassword struct {
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
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
	w.WriteHeader(http.StatusCreated)
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

func Senduser(w http.ResponseWriter, r *http.Request) {
	var user Users
	var returnuser Users
	db, err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err != nil {
		json.NewEncoder(w).Encode(DatabaseError)
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	newid,err:= strconv.ParseUint(id,10,64) // base 10 system 64 bit
	if err!= nil {
		panic(err.Error())
	}
	result := db.Where(&Users{ID:uint(newid)}).First(&user)
	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(DatabaseRetriveError)
	} else {
		returnuser = Users{ID:user.ID,Name:user.Name,Address:user.Address,Phone:user.Phone,Email:user.Email,RRNum:user.RRNum,MeterNo:user.MeterNo,Tariff:user.Tariff,SanctLoad:user.SanctLoad}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(returnuser)
	}
}


func Changepassword(w http.ResponseWriter, r *http.Request) {
	var pass ChangePassword
	var user Users
	db, err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err != nil {
		json.NewEncoder(w).Encode(DatabaseError)
		panic(err.Error())
	}
	defer db.Close()
	
	vars := mux.Vars(r)
	email := vars["email"]
	id := vars["id"]
	newid, err:= strconv.ParseUint(id,10,64)
	if err != nil {
		panic(err.Error())
	}


	_=json.NewDecoder(r.Body).Decode(&pass)
	db.Where(&Users{Email:email, ID:uint(newid)}).First(&user)
	if user.Password != pass.OldPass {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(WrongPassword)
	} else {
		w.Header().Set("Content-Type", "application/json")
		db.Model(&user).Update("password",pass.NewPass)
		json.NewEncoder(w).Encode(PasswordChanged)
	}


}