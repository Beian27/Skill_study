package main

import (
	"fmt"
	"log"
	"net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Web")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)
	
	server := &http.Server{
		Addr:	 ":8080",
		Handler: mux,	
	}

	if err := server.ListenAndServe(); err != nil{
		log.Fatal(err)
	}
}
