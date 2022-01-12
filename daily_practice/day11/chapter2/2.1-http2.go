package main


import(
	"log"
	"net/http"
)

func main(){

	srv := &http.Server{Addr:"127.0.0.1:8080", Handler: http.HandlerFunc(handle)}

	log.Printf("Serving on https:/0.0.0.0:8080")
	log.Fatal(srv.ListenAndServeTLS("day11/chapter2/server.crt","day11/chapter2/server.key"))

}

func handle(w http.ResponseWriter, r *http.Request){
	log.Printf("Got connection: %s", r.Proto)

	w.Write([]byte("Hello this is a HTTP 2 Message!"))
}