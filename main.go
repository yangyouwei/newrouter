package main

import (
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
	workdir := conf.Workdir
	switch {
	case speedmode == "full":
		//停止dnsmasq
		util.Shellout("/etc/init.d/dnsmasq stop",workdir)
		//启动redirect
		util.Shellout("/etc/init.d/redirect start",workdir)
		//加载防火墙
		util.Shellout("",workdir)
	case speedmode == "foreigen":
		//停止dnsmasq
		util.Shellout("/etc/init.d/dnsmasq stop",workdir)
		//启动redirect
		util.Shellout("/etc/init.d/redirect start",workdir)
		//加载防火墙
		util.Shellout("",workdir)
	case speedmode == "multicontry":
		//停止dnsmasq
		util.Shellout("/etc/init.d/dnsmasq stop",workdir)
		//启动redirect
		util.Shellout("/etc/init.d/redirect start",workdir)
		//加载防火墙
		util.Shellout("",workdir)
	case speedmode == "stopspeed":
		//清空防火墙
		util.Shellout("",workdir)
		//关闭redirect
		util.Shellout("/etc/init.d/redirect stop",workdir)
		//重启dnsmasq
		util.Shellout("/etc/init.d/dnsmasq start",workdir)
	default:
		//清空防火墙
		util.Stopspeed()
		//关闭redirect
		util.Shellout("/etc/init.d/redirect stop",workdir)
		//重启dnsmasq
		util.Shellout("/etc/init.d/dnsmasq start",workdir)
	}
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
