package server

import (
	"log"
)

func (r *SQLite) VerifyCredentials(credentials Credentials) bool {
	// Called to verify user credentials with credentials stored to the database
	query, err := r.db.Query("SELECT * FROM forum_login WHERE Mail = ? ", credentials.Email)

	if err != nil {
		log.Fatal(err)
	}

	for query.Next() {
		return true
	}
	return false
}
