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
	connHost = "0.0.0.0"
)

func main() {

	router := mux.NewRouter()
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	router.Handle("/", handlers.LoggingHandler(logFile, cstm_hndlrs.MainPageHandler))
	//router.Handle("/home", handlers.LoggingHandler(logFile, cstm_hndlrs.HomePageHandler))
	router.Handle("/login", handlers.LoggingHandler(logFile, cstm_hndlrs.LoginPageHandler)).Methods("POST", "GET")
	router.Handle("/logout", handlers.LoggingHandler(logFile, cstm_hndlrs.LogoutFormPageHandler)).Methods("POST")
	router.Handle("/books", handlers.LoggingHandler(logFile, cstm_hndlrs.AllBooksHandler))
	router.Handle("/books/reversed", handlers.LoggingHandler(logFile, cstm_hndlrs.RevAllBooksHandler))
	router.Handle("/books/create", handlers.LoggingHandler(logFile, cstm_hndlrs.CreateBooksHandler)).Methods("POST", "GET")
	router.Handle("/books/{id}", handlers.LoggingHandler(logFile, cstm_hndlrs.OneBookHandler))
	router.Handle("/journals", handlers.LoggingHandler(logFile, cstm_hndlrs.AllJournalHandler))
	router.Handle("/journals/reversed", handlers.LoggingHandler(logFile, cstm_hndlrs.RevAllJournalHandler))
	router.Handle("/journals/create", handlers.LoggingHandler(logFile, cstm_hndlrs.CreateJournalHandler)).Methods("POST", "GET")
	router.Handle("/journals/{id}", handlers.LoggingHandler(logFile, cstm_hndlrs.OneJournalHandler))

	err = http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
