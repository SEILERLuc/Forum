package server

import (
	"log"
)

func (r *SQLite) RegisterCredentials(credentials Credentials) bool {
	// Called to store user's credentials
	add, err := r.db.Exec("INSERT INTO forum_login (Mail, Username, Password) VALUES (?,?,?) ", credentials.Email, credentials.Username, credentials.Password)

	if err != nil {
		log.Fatal(err)
	}

	if add == nil {
		return false
	}
	return true
}
