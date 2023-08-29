package server

import (
	"log"
)

func (r *SQLite) GetNameByID(user_id int) string {
	// Called to get a username with the ID
	query, err := r.db.Query("SELECT Username FROM forum_login WHERE ID=?", user_id)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()
	var username string
	for query.Next() {
		err := query.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
	}
	return username
}
