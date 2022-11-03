package main

import (
	"html-project-backend/database"
	rest "html-project-backend/httpd"
)

func main() {
	client, collection := database.Connect()
	println(client)
	//database.Insert(collection)
	//database.Update(collection)
	//database.Find(collection)
	rest.Gin(collection)
}
