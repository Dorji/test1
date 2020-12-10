package main

import (
	"fmt"
	//"github.com/dgrijalva/jwt-go"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"text/template"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	Username string `valid:"alpha,required"`
	Password string `valid:"alpha,required"`
	Age      int
	Phone    string
	Link     string
}

func ValidateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validateError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validateError, "Username")
		passwordError := govalidator.ErrorByField(validateError, "Password")
		if usernameError != "" {
			log.Println("username validation error:", usernameError)
			return valid, "Validation error with username"
		}
		if passwordError != "" {
			log.Println("password validation error:", passwordError)
			return valid, "Validation error with password"
		}

	}
	return valid, "ValidationError"
}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: "NewUser",
		Age:      35,
		Phone:    "+49 999 999 22 33",
		Link:     "github.com/new_user/portfolio",
	}
	parserdTemplate, _ := template.ParseFiles("templates/createbook.html")
	err := parserdTemplate.Execute(w, user)
	if err != nil {
		log.Println("error while parsing template with user:", err)
		return
	}

}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("error while executing template")
			return
		}
	} else {
		user := ReadUserForm(r)
		valid, validationError := ValidateUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, "Error: "+validationError+"!!")
		}
		fmt.Fprintf(w, "Hello "+user.Username+"!!")
	}
}
func ReadUserForm(r *http.Request) *User {
	r.ParseForm()                    //get data from form
	user := new(User)                //default(empty) struct
	decoder := schema.NewDecoder()   // standart decoder
	decoder.Decode(user, r.PostForm) //write into dummy User struct all data from POST
	return user
}

func main() {
	//Реконфигурация static через мультиплексер
	router := mux.NewRouter()

	router.HandleFunc("/login", LoginPageHandler).Methods("GET", "POST")
	//Поддержка самого файл-сервера
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	//Запуск сервера
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}
