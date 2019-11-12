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
	workdir := conf.Workdir
	switch {
	case speedmode == "full":
		fmt.Println("mode is full")
		//停止dnsmasq
		fmt.Println("stop dnsmasq")
		err, standout, standerro := util.Shellout("/etc/init.d/dnsmasq stop",workdir)
		if err != nil {
			fmt.Println("util.shellout error: ",err)
		}else if standout != "" {
			 fmt.Println("exec stand output: "+standout)
		}else if standerro != "" {
			 fmt.Println("exec stand error output "+standerro)
		}
		//启动redirect
		fmt.Println("start redirect")
		err1, standout1, standerro1 := util.Shellout("/etc/init.d/redirect start",workdir)
		if err1 != nil {
			fmt.Println(err1)
		}else if standout1 != "" {
			fmt.Println(standout1)
		}else if standerro1 != "" {
			fmt.Println(standerro1)
		}
		//加载防火墙
		fmt.Println("loading iptables")
		err2, standout2, standerro2 := util.Shellout("",workdir)
		if err2 != nil {
			fmt.Println(err2)
		}else if standout2 != "" {
			fmt.Println(standout2)
		}else if standerro2 != "" {
			fmt.Println(standerro2)
		}
	case speedmode == "foreigen":
		fmt.Println("mode is foreigen")
		//停止dnsmasq
		util.Shellout("/etc/init.d/dnsmasq stop",workdir)
		//启动redirect
		util.Shellout("/etc/init.d/redirect start",workdir)
		//加载防火墙
		util.Shellout("",workdir)
	case speedmode == "multicontry":
		fmt.Println("mode is multicontry")
		//停止dnsmasq
		util.Shellout("/etc/init.d/dnsmasq stop",workdir)
		//启动redirect
		util.Shellout("/etc/init.d/redirect start",workdir)
		//加载防火墙
		util.Shellout("",workdir)
	case speedmode == "stopspeed":
		fmt.Println("mode is stopspeed")
		//清空防火墙
		util.Shellout("",workdir)
		//关闭redirect
		util.Shellout("/etc/init.d/redirect stop",workdir)
		//重启dnsmasq
		util.Shellout("/etc/init.d/dnsmasq start",workdir)
	default:
		fmt.Println("mode is stopspeed")
		//清空防火墙
		util.Stopspeed()
		//关闭redirect
		util.Shellout("/etc/init.d/redirect stop",workdir)
		//重启dnsmasq
		util.Shellout("/etc/init.d/dnsmasq start",workdir)
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
