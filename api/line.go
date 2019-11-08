package api

import (
	"encoding/json"
	"fmt"
	"github.com/yangyouwei/newrouter/conf"
	"github.com/yangyouwei/newrouter/util"
	"log"
	"net/http"
)

type line struct {
	Ipaddr string `json:"ipaddr"`
	Comment string `json:"comment"`
}

type lines struct {
	Lines []line `json:"lines"`
}


func Getlines(w http.ResponseWriter, r *http.Request)  {
	log.Println("request domain ",r.Host,"URL: ",r.URL)
	//读取配置文件
	l := util.ReadAlLFromFile(conf.LinesConfig)
	strline := lines{}
	err := json.Unmarshal([]byte(l),&strline)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strline)
	w.Write([]byte(fmt.Sprint(strline)))
}

func GetUseline(w http.ResponseWriter, r *http.Request)  {

}

func ModLineconf(w http.ResponseWriter, r *http.Request)  {

}

func AppalyLine(w http.ResponseWriter, r *http.Request)  {

}

func SpeedMod(w http.ResponseWriter, r *http.Request)  {
	//全局
	//仅加速国外
	//仅加速指定国家
	//停止加速
}