package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWeb)
	http.HandleFunc("/hi", hiWeb)
	fmt.Println("our app is working")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web!")
}
func hiWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from my web app!")
}
