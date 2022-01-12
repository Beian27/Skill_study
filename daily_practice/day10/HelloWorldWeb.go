package main

import (
	"net/http"
)

type Refer struct {
	handler http.Handler
	refer string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
		//this.handler.ServeHTTP(w, r)
	}
}

func myHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main(){
	//server := &http.Server{
	//	Addr: "0.0.0.0:80",
	//}
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "127.0.0.1",
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":80", referer)
}


