package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"time"
	//jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/mux"
	//"github.com/dgrijalva/jwt-go"
	//_ "github.com/lib/pq"
)

var mySigningKey = []byte("supersecret")

type User struct {
	username string
	password string
}
type Car struct {
	max_speed int
	distance  int
	handler   string
	stock     string
}

var UsersArr = make(map[string]User)
var CarsArr = make(map[string]Car)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/auth",GetToken).Methods("POST")
	router.HandleFunc("/register", RegisterUser).Methods("POST")
	//router.HandleFunc("/auto",GetAuto).Methods("GET")
	//router.HandleFunc("/auto",AddAuto).Methods("POST")
	//router.HandleFunc("/auto",PutAuto).Methods("PUT")
	//router.HandleFunc("/auto",DelAuto).Methods("DELETE")
	router.HandleFunc("/stock", GetAll).Methods("GET")
	fmt.Println("server initialized")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func GetAll(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"Message\" : \"User created. Try to auth\"}")
}
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newUser User
	json.Unmarshal(reqBody, &newUser)

	fmt.Println("try add user")

	if _, ok := UsersArr[newUser.username]; ok {
		fmt.Fprintf(w, "{\"Error\" : \"Auto with that mark exists\"}")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		UsersArr[newUser.username] = newUser
		fmt.Println("user added")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{\"Message\" : \"User created. Try to auth\"}")
	}

}

//var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//		return mySigningKey, nil
//	},
//	SigningMethod: jwt.SigningMethodHS256,
//})
//
//func GetToken(w http.ResponseWriter, r *http.Request){
//
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	var currUser User
//	json.Unmarshal(reqBody, &currUser)
//
//	if currUser.password==UsersArr[currUser.username].username {
//
//
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// Устанавливаем набор параметров для токена
//	claims := make(jwt.MapClaims)
//	claims["admin"] = true
//	claims["name"] = "specialist"
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//	token.Claims = claims
//
//	// Подписываем токен нашим секретным ключем
//	tokenString, _ := token.SignedString(mySigningKey)
//
//	// Отдаем токен клиенту
//	w.Write([]byte(tokenString))
//
//	}
//}
