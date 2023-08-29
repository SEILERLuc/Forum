package server

import (
	"log"
)

func (r *SQLite) GetMessagesByCategory(category string) []WebMessages {
	// Called to get all messages of a category
	rows, err := r.db.Query("SELECT * FROM forum_data_test WHERE Type=?;", category)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var post_messages []WebMessages

	for rows.Next() {
		var tmp_message WebMessages
		if err := rows.Scan(&tmp_message.ID, &tmp_message.UserID, &tmp_message.Username, &tmp_message.Post, &tmp_message.Type); err != nil {
			log.Fatal(err)
		}
		post_messages = append(post_messages, tmp_message)
	}
	return post_messages
}
