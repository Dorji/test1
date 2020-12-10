package main

import (
	cstm_hndlrs "github.com/Dorji/course/day4exercize/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

func main() {

	router := mux.NewRouter()
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	router.Handle("/", handlers.LoggingHandler(logFile, cstm_hndlrs.LoginPageHandler))
	//router.Handle("/home", handlers.LoggingHandler(logFile, cstm_hndlrs.HomePageHandler))
	router.Handle("/login", handlers.LoggingHandler(logFile, cstm_hndlrs.LoginFormPageHandler)).Methods("POST")
	router.Handle("/logout", handlers.LoggingHandler(logFile, cstm_hndlrs.LogoutFormPageHandler)).Methods("POST")
	router.Handle("/books", handlers.LoggingHandler(logFile, cstm_hndlrs.AllBooksHandler))
	router.Handle("/books/create", handlers.LoggingHandler(logFile, cstm_hndlrs.CreateBooksHandler)).Methods("POST")
	router.Handle("/books/create", handlers.LoggingHandler(logFile, cstm_hndlrs.CreateBooksHandler))
	router.Handle("/journals", handlers.LoggingHandler(logFile, cstm_hndlrs.AllJournalHandler))
	router.Handle("/journals/create", handlers.LoggingHandler(logFile, cstm_hndlrs.CreateJournalHandler)).Methods("POST")
	router.Handle("/journals/create", handlers.LoggingHandler(logFile, cstm_hndlrs.CreateJournalHandler))

	err = http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
