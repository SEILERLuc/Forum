package server

import (
	"log"
)

func (r *SQLite) GetPasswordByEmail(email string) string {
	// Called to get a password by an Email
	query, err := r.db.Query("SELECT Password FROM forum_login WHERE Mail=?", email)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()
	var user_pass string
	for query.Next() {
		err := query.Scan(&user_pass)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user_pass
}
