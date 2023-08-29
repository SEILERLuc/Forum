package main

import (
	server "forum/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	server.Serve()
}
