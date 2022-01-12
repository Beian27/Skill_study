package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main () {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/Users/mac/Pictures"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
