package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

type Book struct {
	Title     string
	PageCount int
	Author    string
	Rating    int
}
type User struct {
	UserName string
	Password string
}
type Books struct {
	Books []Book
}

var BookCollection = Books{
	Books: []Book{
		{
			Title:     "Harry Potter",
			Author:    "J.K. Rowling",
			PageCount: 300,
			Rating:    500,
		},
		{
			Title:     "Lord of the Rings",
			Author:    "Tolkien",
			PageCount: 100500,
			Rating:    100500,
		},
		{
			Title:     "The secret of the secrets",
			Author:    "Random",
			PageCount: 100,
			Rating:    100,
		},
	},
}

type Journal struct {
	Edition      string
	SerialNumber int
	PageCount    int
	Redactor     string
}
type Journals struct {
	Journals []Journal
}

var JournalCollection = Journals{
	Journals: []Journal{
		{
			Edition:      "15",
			SerialNumber: 112233,
			PageCount:    3333333,
			Redactor:     "Isinov",
		},
		{
			Edition:      "22",
			SerialNumber: 112233,
			PageCount:    224322,
			Redactor:     "MArvel",
		},
		{
			Edition:      "11",
			SerialNumber: 112233,
			PageCount:    44444444,
			Redactor:     "DC",
		},
	},
}
var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//SetSession ...
func SetSession(userName string, response http.ResponseWriter) {
	value := map[string]string{"username": userName}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

//ClearSession ...
func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//GetUserName ...
func GetUserName(request *http.Request) (userName string) {
	cookie, err := request.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}
