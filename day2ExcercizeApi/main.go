package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	//"time"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("supersecret")

type User struct {
	username string
	password string
}
type Car struct {
	Max_speed string `json:"max_speed"`
	Distance  string `json:"distance"`
	Handler   string `json:"handler"`
	Stock     string `json:"stock"`
}

var UsersArr = make(map[string]User)
var CarsArr = make(map[string]Car)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/auth", GetToken).Methods("POST")
	router.HandleFunc("/register", RegisterUser).Methods("POST")
	router.Handle("/auto/{mark}", jwtMiddleware.Handler(http.HandlerFunc(GetAuto))).Methods("GET")
	router.Handle("/auto/{mark}", jwtMiddleware.Handler(http.HandlerFunc(AddAuto))).Methods("POST")
	router.Handle("/auto/{mark}", jwtMiddleware.Handler(http.HandlerFunc(PutAuto))).Methods("PUT")
	router.Handle("/auto/{mark}", jwtMiddleware.Handler(http.HandlerFunc(DelAuto))).Methods("DELETE")
	router.Handle("/stock", jwtMiddleware.Handler(http.HandlerFunc(GetAll))).Methods("GET")
	fmt.Println("server initialized")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func DelAuto(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	mark := mux.Vars(r)["mark"]
	var newCar Car
	json.Unmarshal(reqBody, &newCar)

	if _, ok := CarsArr[mark]; ok {
		delete(CarsArr, mark)
		fmt.Println("Auto added")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "{\"Message\" : \"Auto deleted\"}")
	} else {
		fmt.Fprintf(w, "{\"Error\" : \"Auto with that mark not found\"}")
		w.WriteHeader(http.StatusNotFound)
	}
}

func AddAuto(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	mark := mux.Vars(r)["mark"]
	var newCar Car
	json.Unmarshal(reqBody, &newCar)

	if _, ok := CarsArr[mark]; ok {
		fmt.Fprintf(w, "{\"Error\" : \"Auto with that mark exists\"}")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		CarsArr[mark] = newCar
		fmt.Println("Auto added")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{\"Message\" : \"Auto created\"}")
	}

}
func PutAuto(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	mark := mux.Vars(r)["mark"]
	var newCar Car
	json.Unmarshal(reqBody, &newCar)

	if _, ok := CarsArr[mark]; ok {
		CarsArr[mark] = newCar
		fmt.Println("Auto added")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{\"Message\" : \"Auto updated\"}")
	} else {
		fmt.Fprintf(w, "{\"Error\" : \"Auto with that mark not found\"}")
		w.WriteHeader(http.StatusBadRequest)
	}

}
func GetAuto(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	mark := mux.Vars(r)["mark"]
	var newCar Car
	json.Unmarshal(reqBody, &newCar)

	if _, ok := CarsArr[mark]; ok {
		resp, err := json.Marshal(CarsArr[mark])
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"Error\" : \"Auto with that mark not found\"}")
	}
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	if len(CarsArr) > 0 {
		resp, err := json.Marshal(CarsArr)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"Error\" : \"No one autos found in DataBase\"}")
	}
}
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newUser User
	json.Unmarshal(reqBody, &newUser)

	fmt.Println("try add user")

	if _, ok := UsersArr[newUser.username]; ok {
		fmt.Fprintf(w, "{\"Error\" : \"User with that name exists\"}")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		UsersArr[newUser.username] = newUser
		fmt.Println("User added")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{\"Message\" : \"User created. Try to auth\"}")
	}

}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func GetToken(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var currUser User
	json.Unmarshal(reqBody, &currUser)

	if currUser.password == UsersArr[currUser.username].password {

		token := jwt.New(jwt.SigningMethodHS256)

		// Устанавливаем набор параметров для токена
		claims := make(jwt.MapClaims)
		claims["admin"] = true
		claims["name"] = "specialist"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		token.Claims = claims

		// Подписываем токен нашим секретным ключем
		tokenString, _ := token.SignedString(mySigningKey)

		// Отдаем токен клиенту
		w.Write([]byte(tokenString))

	}
}
