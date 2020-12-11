package handlers

import (
	"github.com/Dorji/course/day4exercize/cookies"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	//"strconv"
)

//MainPageHandler ...
var MainPageHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			parsedTemplate, _ := template.ParseFiles("templates/main.html")
			parsedTemplate.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	})

//LoginPageHandler ...
var LoginPageHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		//Показываем юзеру форму для логина
		if r.Method == "GET" {
			parsedTemplate, _ := template.ParseFiles("templates/login.html")
			parsedTemplate.Execute(w, nil)
		} else {
			username := r.FormValue("username")
			password := r.FormValue("password")
			target := "/"
			if username != "" && password != "" {
				cookies.SetSession(username, w)
				target = "/"
			}
			http.Redirect(w, r, target, 302)
		}
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

var OneBookHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		one, _ := strconv.Atoi(id)
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			data := cookies.BookCollection.Books[one]
			parsedTemplate, _ := template.ParseFiles("templates/oneBook.html")
			parsedTemplate.Execute(w, data)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	})
var RevAllBooksHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			data := cookies.BookCollection.Books
			sort.SliceStable(data, func(i, j int) bool {
				return data[i].Rating > data[j].Rating
			})
			parsedTemplate, _ := template.ParseFiles("templates/books.html")
			parsedTemplate.Execute(w,
				cookies.Books{
					Books: data,
				})
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
				rating := r.FormValue("rating")
				pc, _ := strconv.Atoi(pagecount)
				rt, _ := strconv.Atoi(rating)
				cookies.BookCollection.Books = append(cookies.BookCollection.Books, cookies.Book{
					Title:     title,
					Author:    author,
					PageCount: pc,
					Rating:    rt,
					Owner:     UserName,
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
var OneJournalHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		one, _ := strconv.Atoi(id)
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			data := cookies.JournalCollection.Journals[one]
			parsedTemplate, _ := template.ParseFiles("templates/oneJournal.html")
			parsedTemplate.Execute(w, data)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	})
var RevAllJournalHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UserName := cookies.GetUserName(r)
		if UserName != "" {
			data := cookies.JournalCollection.Journals
			sort.SliceStable(data, func(i, j int) bool {
				return data[i].PageCount > data[j].PageCount
			})
			parsedTemplate, _ := template.ParseFiles("templates/journals.html")
			parsedTemplate.Execute(w, cookies.Journals{Journals: data})
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
