package server

import (
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// The SIGN IN / SIGN UP page
func LogIn(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	_ = tmpl.Execute(w, nil)
}

// The ERROR page
func ErrorLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/refused.html"))
	_ = tmpl.Execute(w, nil)
}

// The page of POST CREATION
func PostPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/send-message.html"))
	_ = tmpl.Execute(w, nil)
}

// The HOME PAGE : display all posts
func HomePage(w http.ResponseWriter, r *http.Request) {
	show_all_messages := Database.GetAllMessages()
	tmpl := template.Must(template.ParseFiles("index.html"))
	_ = tmpl.Execute(w, struct {
		WebMessages []WebMessages
	}{
		WebMessages: show_all_messages,
	})
}

// The CATEGORY PAGE : display posts of a chosen category
func CategoryPage(w http.ResponseWriter, r *http.Request) {
	filter := r.FormValue("category_filter")
	messages_by_filter := Database.GetMessagesByCategory(filter)
	tmpl := template.Must(template.ParseFiles("index.html"))
	_ = tmpl.Execute(w, struct {
		WebMessages []WebMessages
	}{
		WebMessages: messages_by_filter,
	})
}
