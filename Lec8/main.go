package main

import (
	"log"
	"net/http"
	"os"

	cstm_hndlrs "github.com/Dorji/course/Lec8/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

func main() {
	router := mux.NewRouter()
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	router.Handle("/", handlers.LoggingHandler(logFile, cstm_hndlrs.LoginPageHandler))
	router.Handle("/home", handlers.LoggingHandler(logFile, cstm_hndlrs.HomePageHandler))
	router.Handle("/login", handlers.LoggingHandler(logFile, cstm_hndlrs.LoginFormPageHandler)).Methods("POST")
	router.Handle("/logout", handlers.LoggingHandler(logFile, cstm_hndlrs.LogoutFormPageHandler)).Methods("POST")

	err = http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}
