package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yangyouwei/newrouter/api"
	"github.com/yangyouwei/newrouter/conf"
	"github.com/yangyouwei/newrouter/models"
	"github.com/yangyouwei/newrouter/util"
	"html/template"
	"log"
	"net/http"
)

var system models.Sysstr

func init()  {
	//判断加速模式。如果是不加速，直接退出
	system.GetSYSTEM()
	speedmode := system.SpeedMod
	//fmt.Println(speedmode)
	switch {
	case speedmode == "full":
		fmt.Println("mode is full")
		//停止dnsmasq
		util.SwitchRedirect(true)
		//加载防火墙
		fmt.Println("loading iptables")
		util.ChSpeedMod("fullspeed",util.Port)
	case speedmode == "foreigen":
		fmt.Println("mode is foreigen")
		util.SwitchRedirect(true)
		//加载防火墙
		fmt.Println("loading iptables")
		util.ChSpeedMod("domsticspeed",util.Port)
	case speedmode == "multicontry":
		fmt.Println("mode is multicontry")
		//停止dnsmasq
		util.SwitchRedirect(true)
		//加载防火墙
		fmt.Println("loading iptables")
		util.ChSpeedMod("multispeed",util.Port)
	case speedmode == "stopspeed":
		fmt.Println("mode is stopspeed")
		//清空防火墙
		util.Stopspeed()
		util.SwitchRedirect(false)
	default:
		fmt.Println("mode is stopspeed")
		//清空防火墙
		util.Stopspeed()
		util.SwitchRedirect(false)
	}
	fmt.Println("finish init")
}

func main()  {


	r := mux.NewRouter()
	r.HandleFunc("/", web).Name("index")

	r.PathPrefix("/web").Handler(http.StripPrefix("/web", http.FileServer(http.Dir(conf.StaticPath))))

	s := r.PathPrefix("/line").Subrouter()
	s.HandleFunc("/getlines", api.Getlines)
	s.HandleFunc("/getuseline", api.GetUseline)
	s.HandleFunc("/modline", api.ModLineconf)
	s.HandleFunc("/speedmod", api.SpeedMod)
	s.HandleFunc("/applayline", api.AppalyLine)
	http.ListenAndServe(":3000", r)
}

func web(w http.ResponseWriter, r *http.Request)  {
	log.Println("request domain ",r.Host,"URL: ",r.URL)
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		log.Println("err")
	}
	t.Execute(w, nil)
}
