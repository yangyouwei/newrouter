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
	system.GetSYSTEM()
	util.InitIpset(&system)
	speedmode := system.SpeedMod
	switch {
	case speedmode == "full":
		fmt.Println("mode is full")
		//启动redirect
		util.SwitchRedirect(true)
		//加载防火墙
		fmt.Println("loading iptables")
		util.ChSpeedMod("fullspeed")
	case speedmode == "foreigen":
		fmt.Println("mode is foreigen")
		//启动redirect
		util.SwitchRedirect(true)
		//加载防火墙
		fmt.Println("loading iptables")
		util.ChSpeedMod("domsticspeed")
	case speedmode == "multicontry":
		fmt.Println("mode is multicontry")
		//启动redirect
		util.SwitchRedirect(true)
		//加载防火墙
		fmt.Println("loading iptables")
		util.ChSpeedMod("multispeed")
	case speedmode == "stopspeed":
		fmt.Println("mode is stopspeed")
		//清空防火墙
		util.Stopspeed()
		//启动关闭加速
		util.SwitchRedirect(false)
	default:
		fmt.Println("mode is stopspeed")
		//清空防火墙
		util.Stopspeed()
		//启动关闭加速
		util.SwitchRedirect(false)
	}
	fmt.Println("finish init")
}

func main()  {
	var serverport string = "3000"

	r := mux.NewRouter()
	r.HandleFunc("/", web).Name("index")

	r.PathPrefix("/web").Handler(http.StripPrefix("/web", http.FileServer(http.Dir(conf.StaticPath))))

	s := r.PathPrefix("/line").Subrouter()
	s.HandleFunc("/getlines", api.Getlines)
	s.HandleFunc("/getuseline", api.GetUseline)
	s.HandleFunc("/modline", api.ModLineconf)
	s.HandleFunc("/speedmod", api.SpeedMod)
	s.HandleFunc("/applayline", api.AppalyLine)
	fmt.Println("server is listening on 127.0.0.1:"+serverport)
	http.ListenAndServe(":"+serverport, r)
}

func web(w http.ResponseWriter, r *http.Request)  {
	log.Println("request domain ",r.Host,"URL: ",r.URL)
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		log.Println("err")
	}
	t.Execute(w, nil)
}
