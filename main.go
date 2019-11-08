package main

import (
	"github.com/gorilla/mux"
	"github.com/yangyouwei/newrouter/api"
	"github.com/yangyouwei/newrouter/conf"
	"html/template"
	"log"
	"net/http"
)

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
