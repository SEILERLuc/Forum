package server

import (
	"fmt"
	"log"
)

func (r *SQLite) SendUserMessage(id int, message string, username string, langage string) {
	// Called to store a post into the database
	message_to_send, err := r.db.Exec("INSERT INTO forum_data_test (User_id, User_post, Username, Type) VALUES (?,?,?,?); ", id, message, username, langage)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi")
		log.Fatal(err)
	}
	if message_to_send == nil {
		fmt.Println("Erreur, lors de l'envoi du message")
	} else {
		fmt.Println("Le message a été envoyé")
	}
}
