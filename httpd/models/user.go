package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // Yes I was using a string as the password to test the thing sue me
}
