### 1. Простейший сервер http
```
package main

import (
"fmt"
"log"
"net/http"
)

const (
CONN_PORT = "8080"
CONN_HOST = "localhost"
)

func helloWorld(w http.ResponseWriter, r *http.Request)  { fmt.Fprintf(w, "Hello World!!!")
}

func main(){ // определить соотношение вида url => handleFunc(функци отображения)
http.HandleFunc("/",helloWorld)
err:=http.ListenAndServe(CONN_HOST+":"+CONN_PORT,nil)
if err!=nil{ log.Fatal("error starting server:",err)
return } }
```
### 2. Простейшая аутентификация

***Аутентификация*** процесс узнавания свой-чужой
***AuthWrapper***
```
func AuthWrapper(handler http.HandlerFunc, realm string) http.HandlerFunc { return func(w http.ResponseWriter, r *
http.Request) { user, pass, ok := r.BasicAuth() // возвращает всю инфу о текущем пользователе if !ok ||
subtle.ConstantTimeCompare([]byte(user), []byte(USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(
PASSWORD)) != 1 { w.Header().Set("www-my-auth-service",`Basic realm="`+realm+`"`)
w.WriteHeader(401)
w.Write([]byte("You unauthorized. to Access"))
} handler(w,r)
} }
```
### 3. gzip
```func main() {
	// определим базовый рутер  multiplexer
	mux:=http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	//compress  преобразует все в gzip
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
```
***Логирование  by gorilla***
```
func main() {
router:=mux.NewRouter()
logFile,err:=os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)

	router.Handle("/",handlers.LoggingHandler(os.Stdout,custom_handlers.GetRequestHandler)).Methods("GET")
	router.Handle("/post",handlers.LoggingHandler(logFile,custom_handlers.PostRequestHandler)).Methods("POST")
	router.Handle("/hello/{name}",handlers.LoggingHandler(os.Stdout,custom_handlers.GetPutRequestHandler)).Methods("GET","PUT")

	err = http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}```




