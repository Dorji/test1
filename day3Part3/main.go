package main

import (
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

var cookieHandler *securecookie.SecureCookie

func init() {
	cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}
func createCookie(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{"username": "Alex"}
	base64Encoded, err := cookieHandler.Encode(key, value)
	if err != nil {
		cookie := &http.Cookie{
			Name:  "fiessrs",
			Value: base64Encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	w.Write([]byte("Cookie created!"))
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	log.Println("Now readin cookie")
	cookie, err := r.Cookie("first-cookie")
	if cookie != nil && err == nil {
		//value:=make(map[string]string)

	}
}

func main() {
	http.HandleFunc("/create", createCookie)
	http.HandleFunc("/read", readCookie)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
