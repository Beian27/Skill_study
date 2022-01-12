package main

import (
	"fmt"
	"log"
	"net/http"
)

type WelcomeHandler struct {
	Language	string
}

func (h WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(h.Language                            ))
	fmt.Println(w, "%s", h.Language)
}

func main(){
	mux := http.NewServeMux()
	mux.Handle("/cn", WelcomeHandler{
		Language: "Beian_27",
	})

	mux.Handle("/en", WelcomeHandler{
		Language: "Welcome Beijing!!!",
	})

	server := &http.Server {
		Addr:		":8080",
		Handler:	mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}


