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

var System models.Sysstr

func init()  {
	System.GetSYSTEM()
	util.RstoreIpset(&System)
	util.SpeedCtl(System.SpeedMod)
	fmt.Println("finish init")
}

func main()  {
	var serverport string = "80"

	r := mux.NewRouter()
	r.HandleFunc("/", web).Name("index")

	r.PathPrefix("/web").Handler(http.StripPrefix("/web", http.FileServer(http.Dir(conf.StaticPath))))

	s := r.PathPrefix("/line").Subrouter()
	s.HandleFunc("/getlines", api.Getlines)
	s.HandleFunc("/getuseline", api.GetUseline)
	s.HandleFunc("/modline", api.ModLineconf)
	s.HandleFunc("/speedmod", api.SpeedMod)
	s.HandleFunc("/applayline", api.AppalyLine)

	sys := r.PathPrefix("/sys").Subrouter()
	sys.HandleFunc("/waninfo", api.GetWanInfo)
	sys.HandleFunc("/restart", api.Rebootsys)

	wifi := r.PathPrefix("/wifi").Subrouter()
	wifi.HandleFunc("/ssid", api.GetWifiInfo)
	wifi.HandleFunc("/setwifi", api.SetWifi)
	wifi.HandleFunc("/restartwifi", api.RestartWifi)

	blacklist := r.PathPrefix("/blacklist").Subrouter()
	blacklist.HandleFunc("/getblack", api.GetBlacklist)
	blacklist.HandleFunc("/setblack", api.SetBlacklist)

	userhosts := r.PathPrefix("/hosts").Subrouter()
	userhosts.HandleFunc("/gethosts", api.GetHosts)
	userhosts.HandleFunc("/sethosts", api.SetHosts)

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
