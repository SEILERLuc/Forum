package server

// Structures we need for the project

// To store a user
type Credentials struct {
	id       int
	Email    string
	Username string
	Password string
}

// To store a message of a user
type WebMessages struct {
	ID       int
	UserID   int
	Username string
	Post     string
	Type     string
}
