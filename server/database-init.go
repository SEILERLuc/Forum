package server

import (
	"database/sql"
)

type SQLite struct { // New type for the database
	db *sql.DB
}

var Database *SQLite

func Init(db *sql.DB) *SQLite { // Init of a database
	Database = &SQLite{
		db: db,
	}
	return Database
}

func (r *SQLite) MigrateLogin() error { // Creation of LOGIN TABLE of not exists
	query := `CREATE TABLE forum_login (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Mail TEXT UNIQUE,
		Username TEXT UNIQUE,
		Password TEXT
	);`
	_, err := r.db.Query(query)
	return err
}

func (r *SQLite) MigrateData() error { // Creation of DATA TABLE if not exists
	query := `CREATE TABLE forum_data_test (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		User_id INTEGER,
		Username TEXT,
		User_post TEXT,
		Type TEXT,
		FOREIGN KEY (User_id) REFERENCES forum_login (ID),
		FOREIGN KEY (Username) REFERENCES forum_login (Username)
	);`
	_, err := r.db.Query(query)
	return err
}
