package models

import (
	"fmt"
	"github.com/yangyouwei/newrouter/db"
)

type Sysstr struct {
	Contry string	`json:"contry"`
	SpeedMod string `json:"speedmod"`
	GetchontryUrl string `json:"getcontryurl"`
	GetLineUrl string `json:"getlineurl"`
}


func (system *Sysstr)GetSYSTEM(){
	err := db.SqlDB.QueryRow("SELECT * FROM system WHERE rowid=?", 1).Scan(&system.Contry, &system.SpeedMod,&system.GetchontryUrl,&system.GetLineUrl)
	if err != nil {
		fmt.Println(err)
	}
}
