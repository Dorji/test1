package handlers

import "net/http"

// функция-отображение в виде промежуточной переменной
var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world from getRequestHandler with variable creation!"))
	})

var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world from PostRequestHandler with variable creation!"))
	})

var GetPutRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world from GetPutRequestHandler with variable creation!"))
	})
