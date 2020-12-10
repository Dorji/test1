package handlers

import (
	"github.com/Dorji/course/day4exercize/cookies"
	"html/template"
	"net/http"
	"strconv"
	//"strconv"
)

//LoginPageHandler ...
var LoginPageHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		//Показываем юзеру форму для логина
		parsedTemplate, _ := template.ParseFiles("templates/login.html")
		parsedTemplate.Execute(w, nil)
	})

//HomePageHandler ...
var HomePageHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		//Извлечем имя пользователя из сессии и подадим на вход шаблону createbook.html
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			data := map[string]interface{}{
				"UserName": UserName,
			}
			parsedTemplate, _ := template.ParseFiles("templates/createbook.html")
			parsedTemplate.Execute(w, data)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	})

//LoginFormPageHandler ...

var LoginFormPageHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		target := "/"
		if username != "" && password != "" {
			cookies.SetSession(username, w)
			target = "/books"
		}
		http.Redirect(w, r, target, 302)
	})

//LogoutFormPageHandler ...

var LogoutFormPageHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		//Очищаем сессию и редиректим на LoginPage
		cookies.ClearSession(w)
		http.Redirect(w, r, "/", 302)
	})

var AllBooksHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			data := cookies.BookCollection
			parsedTemplate, _ := template.ParseFiles("templates/books.html")
			parsedTemplate.Execute(w, data)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	})
var CreateBooksHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			if r.Method == "POST" {

				title := r.FormValue("title")
				author := r.FormValue("author")
				pagecount := r.FormValue("pagecount")
				pc, _ := strconv.Atoi(pagecount)
				cookies.BookCollection.Books = append(cookies.BookCollection.Books, cookies.Book{
					Title:     title,
					Author:    author,
					PageCount: pc,
				})
				http.Redirect(w, r, "/books", 302)
			} else {
				parsedTemplate, _ := template.ParseFiles("templates/createbook.html")
				parsedTemplate.Execute(w, nil)
			}
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	})
var AllJournalHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			data := cookies.JournalCollection
			parsedTemplate, _ := template.ParseFiles("templates/journals.html")
			parsedTemplate.Execute(w, data)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	})
var CreateJournalHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			if r.Method == "POST" {

				serialnumber := r.FormValue("serialnumber")
				edition := r.FormValue("edition")
				pagecount := r.FormValue("pagecount")
				redactor := r.FormValue("redactor")
				pc, _ := strconv.Atoi(pagecount)
				sn, _ := strconv.Atoi(serialnumber)
				cookies.JournalCollection.Journals = append(cookies.JournalCollection.Journals, cookies.Journal{
					SerialNumber: sn,
					PageCount:    pc,
					Edition:      edition,
					Redactor:     redactor,
				})
				http.Redirect(w, r, "/journals", 302)
			} else {
				parsedTemplate, _ := template.ParseFiles("templates/createJournal.html")
				parsedTemplate.Execute(w, nil)
			}
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	})
