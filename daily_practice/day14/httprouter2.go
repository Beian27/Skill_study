package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default get!!!"))
	})

	router.POST("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default post!!!"))
	})

	//router.GET("/user/hello", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//	w.Write([]byte("user name:" + p.ByName("name")))
	//})

	router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name:" + p.ByName("name")))
		w.Write([]byte(p.ByName("name")))
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
