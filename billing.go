package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"encoding/json"
	"math/rand"
)

type Fixed struct {
	FirstKW    string `json:"first_kw"`     // 1 KW  Rs.40
	EveryAddKW string `json:"every_add_kw"` // After every one KW Rs.50
}

type Variable struct {
	Slab1 string `json:"slab_1"` // Slab 1 (0-30 Units) Rs. 3.25/Unit
	Slab2 string `json:"slab_2"` // Slab 2 (31-100 Units) Rs. 4.70/Unit
	Slab3 string `json:"slab_3"` // Slab 3 (101-200 Units) Rs. 6.25/Unit
	Slab4 string `json:"slab_4"` // Slab 4 (201-300 Units) Rs. 7.30/Unit
	Slab5 string `json:"slab_5"` // Slab 5 (301-400 Units) Rs. 7.35/Unit
	Slab6 string `json:"slab_6"` // Slab 6 (Above 400 Units) Rs. 7.40/Unit
}

type Bill struct {
	RRNum             string   `json:"rr_num"`     					// Revenue Registration Number ( Unique ID )
	AccountID         uint64   `json:"account_id"` 					// Account ID
	Email             string   `json:"email"`						// Email
	Phone             uint64   `json:"phone"`						// Phone
	MtrReadCode       string   `json:"mtr_read_code"`               // Meter Read Code
	Name              string   `json:"name"`                        // Name of Register User
	Address           string   `json:"address"`                     // Address of User
	Tariff            string   `json:"tariff"`                      // Tarrif BBMP, Pancayat etc
	SanctLoad         string   `json:"sanct_load"`                  // Sanctioned Load HP and KW
	FromDate          string   `json:"from_date"`                   // From Date
	ToDate            string   `json:"to_date"`                     // To Date
	ReadingDate       string   `json:"reading_date"`                // Reading Date
	BillNum           uint32   `gorm:"primary_key" json:"bill_num"` // Bill Number
	MeterSLNum        uint64   `json:"meter_sl_num"`                // Meter Serial Number
	PresentRead       string   `json:"present_read"`                // Present Load
	PreviousRead      string   `json:"previous_read"`               // Previous Load
	Constant          string   `json:"constant"`                    // Constant
	Consumption       string   `json:"consumption"`                 // Consumption ( Present - Previous )
	Average           string   `json:"average"`                     // Average
	RecordedMD        string   `json:"recorded_md"`                 // Recorded MD
	PowerFactor       string   `json:"power_factor"`                // Power Factor
	FixedCharges      Fixed    `json:"fixed_charges"`               // Fixed charges
	VariableCharges   Variable `json:"variable_charges"`            // Variable Charges
	RebatesTODCharges string   `json:"rebates_tod_charges"`         // Rebates, TOD Charges
	PFPenalty         string   `json:"pf_penalty"`                  // Power Factor Penalty
	ExLoadMDPenalty   string   `json:"ex_load_md_penalty"`          // Extra Load Penalty
	Interest          string   `json:"interest"`                    // Interest
	Others            string   `json:"others"`                      // Others
	Tax               string   `json:"tax"`                         // Tax GST CGST+SGST
	CurrentBillAmt    float64  `json:"current_bill_amt"`            // Current Bill Amount
	Arrears           float64  `json:"arrears"`                     // Arrears
	CreditsAdj        string   `json:"credits_adj"`                 // Credits & Adj
	GovSubsidy        float64  `json:"gov_subsidy"`                 // Government Subsidy
	NetAmtDue         float64  `json:"net_amt_due"`                 // Net Amount to be Paid
	DueDate           string   `json:"due_date"`                    // Due Date
	Paid              bool     `json:"paid"`
}

type MeterResponse struct {
	RRNum string `json:"rr_num"` // RR Number
}

// var random uint32

func CreateBill(w http.ResponseWriter, r *http.Request) {
	//var meterresponse MeterResponse
	//random=0
	var bill Bill
	//var findbill Bill
	var user Users
	var kwhupdate KWHUpdate
	// json.NewDecoder(r.Body).Decode(&meterresponse)

	vars := mux.Vars(r)
	rrnum := vars["rr_num"]

	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Bill{})
	Showlog(w, r)



	db.Where(&KWHUpdate{RRNum: rrnum}).First(&kwhupdate)
	db.Where(&Users{RRNum: rrnum}).First(&user)

	fmt.Println(kwhupdate)
	fmt.Println(user)
	// create bill
	data := strconv.FormatFloat(kwhupdate.KWH, 'f', 10, 64)

	t := time.Now().Local()
	duedate := time.Now().AddDate(0, 1, 0)

	fromdate := time.Now().AddDate(0, -1, 0)

	billamount := (kwhupdate.KWH * 5)

	random := rand.Uint32()
	mydata := strconv.FormatFloat(billamount, 'f', 7, 64)
	urdata, err := strconv.ParseFloat(mydata, 8)
	if err != nil {
		panic(err)
	}
	bill = Bill{RRNum: user.RRNum, AccountID: user.ID, Name: user.Name, Address: user.Address, Consumption: data, MeterSLNum: user.MeterNo, SanctLoad: user.SanctLoad, Tariff: user.Tariff, ReadingDate: t.Format("2006-01-02"), CurrentBillAmt: billamount, Arrears: 0, GovSubsidy: 0, NetAmtDue: urdata, DueDate: duedate.Format("2006-01-02"), FromDate: fromdate.Format("2006-01-02"), ToDate: t.Format("2006-01-02"), BillNum: random, Paid: false, Email:user.Email, Phone:user.Phone}

	//db.Where(&Bill{RRNum: meterresponse.RRNum}).First(&findbill)
	//if findbill.RRNum != kwhupdate.RRNum {
	db.Create(&bill)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Billgenerated)
	//} else {
	//	db.Model(&bill).Update("consumption", kwhupdate.KWH)
	//}

}

func PaidBill(w http.ResponseWriter, r *http.Request) {
	Showlog(w, r)
	var bill Bill
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	rrnum := vars["rr_num"]
	billnum := vars["bill_num"]

	data, err := strconv.ParseFloat(billnum, 64)
	if err != nil {
		panic(err)
	}
	db.Where(&Bill{RRNum: rrnum, BillNum: uint32(data)}).Find(&bill)
	//db.Where(&Bill{RRNum:rrnum, BillNum:uint32(data)}).First(&bill)

	db.Model(&bill).Update("paid", true)
	fmt.Println(bill)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PaidSuccess)

}

func UnpaidBill(w http.ResponseWriter, r *http.Request) {
	Showlog(w, r)
	var bill Bill
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	rrnum := vars["rr_num"]
	billnum := vars["bill_num"]

	data, err := strconv.ParseFloat(billnum, 64)
	if err != nil {
		panic(err)
	}
	db.Where(&Bill{RRNum: rrnum, BillNum: uint32(data)}).Find(&bill)
	//db.Where(&Bill{RRNum:rrnum, BillNum:uint32(data)}).First(&bill)

	db.Model(&bill).Update("paid", false)
	fmt.Println(bill)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UnpaidSuccess)
}

func GetBills(w http.ResponseWriter, r *http.Request) {
	var bills []Bill

	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	vars := mux.Vars(r)
	rrnum := vars["rr_num"]

	db.Where(&Bill{RRNum: rrnum}).Find(&bills)
	Showlog(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bills)

	//fmt.Println(bills)

}

// main API
func CreateAllBill(w http.ResponseWriter, r *http.Request) {
	var user []Users
	var kwhupdate KWHUpdate
	var bill Bill
	blank := KWHUpdate{}
	db, err := gorm.Open("mysql", "session:session@/project2018")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Bill{})
	Showlog(w, r)



	//db.Where(&KWHUpdate{RRNum: rrnum}).First(&kwhupdate)
	db.Where(&Users{}).Find(&user)
	fmt.Println(len(user))
	for i:=0;i<len(user);i++  {
		kwhupdate = KWHUpdate{}
		db.Where(&KWHUpdate{RRNum: user[i].RRNum}).Find(&kwhupdate)
		if kwhupdate != blank {

			// create bill
			data := strconv.FormatFloat(kwhupdate.KWH, 'f', 10, 64)

			t := time.Now().Local()
			duedate := time.Now().AddDate(0, 1, 0)

			fromdate := time.Now().AddDate(0, -1, 0)

			billamount := (kwhupdate.KWH * 5)

			random := rand.Uint32()
			mydata := strconv.FormatFloat(billamount, 'f', 7, 64)
			urdata, err := strconv.ParseFloat(mydata, 8)
			if err != nil {
				panic(err)
			}
			bill = Bill{RRNum: user[i].RRNum, AccountID: user[i].ID, Name: user[i].Name, Address: user[i].Address, Consumption: data, MeterSLNum: user[i].MeterNo, SanctLoad: user[i].SanctLoad, Tariff: user[i].Tariff, ReadingDate: t.Format("2006-01-02"), CurrentBillAmt: billamount, Arrears: 0, GovSubsidy: 0, NetAmtDue: urdata, DueDate: duedate.Format("2006-01-02"), FromDate: fromdate.Format("2006-01-02"), ToDate: t.Format("2006-01-02"), BillNum: random, Paid: false, Email:user[i].Email, Phone:user[i].Phone}

			//db.Where(&Bill{RRNum: meterresponse.RRNum}).First(&findbill)
			//if findbill.RRNum != kwhupdate.RRNum {
			db.Create(&bill)
			fmt.Println("Bill created for "+ user[i].Email)

		}
	}


}