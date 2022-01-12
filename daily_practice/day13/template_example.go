package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Userinfo struct {
	Name 	string
	Gender	string
	Age		int
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./day13/template_example.html")
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}
	userinfo := Userinfo{
		Name:   "李四",
		Gender: "男",
		Age:    1,
	}
	t.Execute(w, userinfo)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
