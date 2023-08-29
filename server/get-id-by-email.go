package server

import (
	"log"
)

func (r *SQLite) GetIDbyEmail(email string) int {
	// Called to get an ID with the Email
	query, err := r.db.Query("SELECT ID FROM forum_login WHERE Mail=?", email)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()
	var id int
	for query.Next() {
		err := query.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return id
}
