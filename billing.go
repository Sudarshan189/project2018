package main

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"encoding/json"
)

type BillPeriod struct {
	FromDate string `json:"from_date"`		// From Date
	ToDate   string `json:"to_date"`		// To Date
}


type Fixed struct {
	FirstKW    string `json:"first_kw"`			// 1 KW  Rs.40
	EveryAddKW string `json:"every_add_kw"`		// After every one KW Rs.50
}

type Variable struct {
	Slab1 string `json:"slab_1"`		// Slab 1 (0-30 Units) Rs. 3.25/Unit
	Slab2 string `json:"slab_2"`		// Slab 2 (31-100 Units) Rs. 4.70/Unit
	Slab3 string `json:"slab_3"`		// Slab 3 (101-200 Units) Rs. 6.25/Unit
	Slab4 string `json:"slab_4"`		// Slab 4 (201-300 Units) Rs. 7.30/Unit
	Slab5 string `json:"slab_5"`		// Slab 5 (301-400 Units) Rs. 7.35/Unit
	Slab6 string `json:"slab_6"`		// Slab 6 (Above 400 Units) Rs. 7.40/Unit
}


type Bill struct {

	RRNum 				string 		`gorm:"primary_key" json:"rr_num"`					// Revenue Registration Number ( Unique ID )
	AccountID 			uint 		`gorm:"unique" json:"account_id"`				// Account ID
	MtrReadCode 		string 		`json:"mtr_read_code"`			// Meter Read Code
	Name 				string 		`json:"name"`					// Name of Register User
	Address 			string 		`json:"address"`				// Address of User
	Tariff 				string 		`json:"tariff"`					// Tarrif BBMP, Pancayat etc
	SanctLoad 			string 		`json:"sanct_load"`				// Sanctioned Load HP and KW
	BillingPeriod 		BillPeriod 	`json:"billing_period"`			// Billing Period From - To
	ReadingDate 		string 		`json:"reading_date"`			// Reading Date
	BillNum 			string 		`json:"bill_num"`				// Bill Number
	MeterSLNum 			uint 		`json:"meter_sl_num"`			// Meter Serial Number
	PresentRead 		string 		`json:"present_read"`			// Present Load
	PreviousRead 		string 		`json:"previous_read"`			// Previous Load
	Constant 			string 		`json:"constant"`				// Constant
	Consumption 		string 		`json:"consumption"`			// Consumption ( Present - Previous )
	Average 			string 		`json:"average"`			 	// Average
	RecordedMD 			string 		`json:"recorded_md"`			// Recorded MD
	PowerFactor 		string 		`json:"power_factor"`			// Power Factor
	FixedCharges 		Fixed 		`json:"fixed_charges"`			// Fixed charges
	VariableCharges 	Variable 	`json:"variable_charges"`		// Variable Charges
	RebatesTODCharges 	string 		`json:"rebates_tod_charges"`    // Rebates, TOD Charges
	PFPenalty 			string 		`json:"pf_penalty"`				// Power Factor Penalty
	ExLoadMDPenalty 	string 		`json:"ex_load_md_penalty"`		// Extra Load Penalty
	Interest 			string 		`json:"interest"`				// Interest
	Others 				string 		`json:"others"`					// Others
	Tax 				string 		`json:"tax"`					// Tax GST CGST+SGST
	CurrentBillAmt 		string 		`json:"current_bill_amt"`		// Current Bill Amount
	Arrears 			string 		`json:"arrears"`				// Arrears
	CreditsAdj 			string 		`json:"credits_adj"`			// Credits & Adj
	GovSubsidy 			string 		`json:"gov_subsidy"`			// Government Subsidy
	NetAmtDue 			string 		`json:"net_amt_due"`			// Net Amount to be Paid
	DueDate 			string 		`json:"due_date"`				// Due Date
}

type MeterResponse struct {
	RRNum     string 	`json:"rr_num"`	// RR Number
}


func CreateBill(w http.ResponseWriter, r *http.Request) {
	var meterresponse MeterResponse
	var bill Bill
	var findbill Bill
	var user Users
	var kwhupdate KWHUpdate
	json.NewDecoder(r.Body).Decode(&meterresponse)
	db,err := gorm.Open("mysql", "akshay:deepika019@/project2018")
	if err!= nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Bill{})
	Showlog(w,r)
	db.Where(&KWHUpdate{RRNum:meterresponse.RRNum}).First(&kwhupdate)
	db.Where(&Users{RRNum:meterresponse.RRNum}).First(&user)
	// create bill
	bill = Bill{RRNum:kwhupdate.RRNum, AccountID:user.ID, Name:user.Name, Address:user.Address, Consumption:kwhupdate.KWH, MeterSLNum:user.MeterNo, SanctLoad:user.SanctLoad, Tariff:user.Tariff}

	db.Where(&Bill{RRNum:meterresponse.RRNum}).First(&findbill)
	 if findbill.RRNum != kwhupdate.RRNum {
		db.Create(&bill)
	 } else {
		 db.Model(&bill).Update("consumption", kwhupdate.KWH)
	 }

}


