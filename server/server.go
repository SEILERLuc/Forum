package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var credentials Credentials // credential type for the user informations
var messages WebMessages    // message type to send the post on a user

func Serve() { // the function that create the server and the URL to display the data
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	Database := Init(db)                            // Initialize the database in database.go
	if err := Database.MigrateLogin(); err != nil { // create the tables if not exist
		fmt.Println(err)
	}
	if err := Database.MigrateData(); err != nil {
		fmt.Println(err)
	}
	// Create and initialise the server, and differents routes
	server := http.NewServeMux()
	server.HandleFunc("/", HomePage)
	server.HandleFunc("/create", Create)
	server.HandleFunc("/login", Register)
	server.HandleFunc("/send", Send)
	server.HandleFunc("/log", LogIn)
	server.HandleFunc("/message", PostPage)
	server.HandleFunc("/category", CategoryPage)
	server.HandleFunc("/error_login", ErrorLogin)
	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("Server listening on http://localhost:8000")
	http.ListenAndServe(":8000", server)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// this function is used for the REGISTRATION OF USERS'S credentials
	if r.Method == "POST" {
		credentials.Email = r.FormValue("email")
		credentials.Username = r.FormValue("username")
		credentials.Password = r.FormValue("password")
		fmt.Println(credentials)
		hashed_user_password := Database.GetPasswordByEmail(credentials.Email)
		match := CheckPasswordHash(credentials.Password, hashed_user_password)
		fmt.Println(credentials)
		isRegister := Database.VerifyCredentials(credentials)
		if isRegister && match {
			http.Redirect(w, r, "http://localhost:8000/", 302)
			HomePage(w, r)
		} else {
			http.Redirect(w, r, "http://localhost:8000/error_login", 302)
			ErrorLogin(w, r)
		}
	} else {
		http.Redirect(w, r, "http://localhost:8000/", 302)
		HomePage(w, r)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	// this function is used for the CREATION AND STORAGE OF USERS'S credentials
	credentials.Email = r.FormValue("email")
	credentials.Username = r.FormValue("username")
	credentials.Password = r.FormValue("password")
	credentials.Password, _ = HashPassword(credentials.Password)
	isCreated := Database.RegisterCredentials(credentials)
	if !isCreated {
		http.Redirect(w, r, "http://localhost:8000/error_login", 302)
		ErrorLogin(w, r)
	} else {
		http.Redirect(w, r, "http://localhost:8000/log", 302)
		LogIn(w, r)
	}
}

func Send(w http.ResponseWriter, r *http.Request) {
	// this function is used to SEND POST / MESSAGE OF USERS
	if r.Method == "POST" {
		user_message := r.FormValue("user_message")
		user_langage := r.FormValue("langage")
		credentials.id = Database.GetIDbyEmail(credentials.Email)
		credentials.Username = Database.GetNameByID(credentials.id)
		if credentials.id == 0 { // if the user is not log in on the forum
			http.Redirect(w, r, "http://localhost:8000/log", 302)
			LogIn(w, r)
		} else {
			Database.SendUserMessage(credentials.id, user_message, credentials.Username, user_langage)
			http.Redirect(w, r, "http://localhost:8000/", 302)
			HomePage(w, r)
		}
	} else {
		http.Redirect(w, r, "http://localhost:8000/", 302)
		HomePage(w, r)
	}
}
