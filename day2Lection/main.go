package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	CONN_PORT = "8083"
	CONN_HOST = "localhost"
)

type User struct {
	Age  int
	Name string
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Age: 35,
		Name: "Bob"}
	parsedTemplate, _ := template.ParseFiles("templates/home.html")
	err := parsedTemplate.Execute(w, user)
	if err != nil {
		log.Printf("template exec error")
		return
	}
}
func main() {

	http.HandleFunc("/", homePageHandler)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
