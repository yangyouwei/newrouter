package api

import (
	"github.com/yangyouwei/newrouter/conf"
	"github.com/yangyouwei/newrouter/util"
	"log"
	"net/http"
)

func Getlines(w http.ResponseWriter, r *http.Request)  {
	log.Println("request domain ",r.Host,"URL: ",r.URL)
	//读取配置文件
	l := util.ReadAlLFromFile(conf.LinesConfig)
	w.Write([]byte(l))
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