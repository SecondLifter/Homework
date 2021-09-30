package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/header",indexHandler)
	mux.HandleFunc("/version",versionHandler)
	mux.HandleFunc("/log",logHandler)
	mux.HandleFunc("/heathz",heathz)
	http.ListenAndServe(":8080",mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	log.Print(r.Header)
	for k,v :=range r.Header{
		vv :=fmt.Sprintf(strings.Join(v,""))
		w.Header().Add(k,vv)
		log.Println(w.Header())
	}

}

func versionHandler(w http.ResponseWriter,r *http.Request)  {
	version :=os.Getenv("VERSION")
	w.Header().Add("Version",version)
	log.Println(w.Header())
}

func logHandler(w http.ResponseWriter,r *http.Request)  {
	IP := r.Header.Get("X-Real-Ip")
	if IP ==""{
		IP =r.Header.Get("X-Forwarded-For")
	}
	if IP == ""{
		IP =r.RemoteAddr
	}

	log.Println(IP,http.StatusOK)
}

func heathz(w http.ResponseWriter,r *http.Request)  {
	 w.WriteHeader(http.StatusOK)

}