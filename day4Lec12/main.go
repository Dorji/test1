package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	CONN_HOST      = "localhost"
	CONN_PORT      = "8080"
	driverName     = "mysql"
	datasourceName = "root:password@/mydb"
)

var db *sql.DB
var connectionError error

type Employee struct {
	ID   uint   `json:"uid"`
	name string `json:"name"`
}

func init() {
	db, connectionError := sql.Open(driverName, datasourceName)
	if connectionError != nil {
		log.Fatal("error while connecting to database:", connectionError)
	}

}
func GetCurrentDB(w http.ResponseWriter, r *http.Request) {

}
func ReaAllEmployees(w http.ResponseWriter, r *http.Request) {
	var query = "SELECT * FROM employee"

}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Println("ready to update name ...")
	}

}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Println("ready to update name ...")
	}

}

func main() {
	router := mux.NewRouter()
	defer db.Close()
	router.HandleFunc("/", GetCurrentDB)
	router.HandleFunc("/employee/create", GetCurrentDB)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
