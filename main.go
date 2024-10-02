package main

import (
	"UrlShorteningService/controllers"
	"UrlShorteningService/database"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Url Shorterning service started....")

	//connect Mongo
	database.ConnectMongoDB()

	//handlers
	http.HandleFunc("POST /shorten", controllers.CreateShorten)

}
