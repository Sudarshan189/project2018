package main
import (
	"net/http"
	"encoding/json"
	"time"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Socket struct {
	RRNum     string `json:"rr_num"`
	SocketId  string `gorm:"primary_key" json:"socket_id"`
	SocStatus string   `json:"soc_status"`
	Limit 		string `gorm:"default:0 " json:"limit"`
	UpdatedAt *time.Time `json:"updated_at"`
}



func SharpThing(w http.ResponseWriter, r *http.Request) {
	var soc1 Socket
	var emp Socket = Socket{}
	var found Socket
	var found1 []Socket = []Socket{}

	db,err := gorm.Open("mysql", "session:session@/project2018")
	if err!= nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Socket{})
	Showlog(w, r)
	// _=json.NewDecoder(r.Body).Decode(&socket)
	vars := mux.Vars(r)
	rrnum := vars["rr_num"]
	socno := vars["soc_num"]
	sta := vars["sta"]
	// fmt.Println(sta)
	//
	// fmt.Println(socket)
	soc1= Socket{RRNum:rrnum,SocketId:socno,SocStatus:sta}

	db.Where(&Socket{RRNum:soc1.RRNum, SocketId:soc1.SocketId}).First(&found)

	if found == emp {
		db.Create(&soc1)
	} else {
		db.Model(&soc1).Update("soc_status", sta)
	}
	//fmt.Println(soc1)
	w.Header().Set("content-type", "application/json")
	db.Where(&Socket{RRNum:soc1.RRNum}).Find(&found1)
	json.NewEncoder(w).Encode(found1)

	// db.Where(&Socket{RRNum:socket.RRNum, SocketId:socket.SocketId}).First()


}

func GetSwitchState(w http.ResponseWriter, r *http.Request) {
 	var soc1 Socket
	// var emp Socket = Socket{}
	var found []Socket = []Socket{}

	db,err := gorm.Open("mysql", "session:session@/project2018")
	if err!= nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Socket{})
	Showlog(w, r)
	// _=json.NewDecoder(r.Body).Decode(&socket)
	vars := mux.Vars(r)
	rrnum := vars["rr_num"]
	// fmt.Println(sta)
	//
	// fmt.Println(socket)
	 soc1= Socket{RRNum:rrnum}

	db.Where(&Socket{RRNum:soc1.RRNum}).Find(&found)
	 //fmt.Println(found[0].UpdatedAt)

	//fmt.Println(soc1)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(found)



}
