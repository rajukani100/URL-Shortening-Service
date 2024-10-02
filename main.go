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

	router := http.NewServeMux()

	//handlers
	router.HandleFunc("POST /shorten", controllers.CreateShorten)
	router.HandleFunc("GET /shorten/{shortenCode}", controllers.RetrieveShorten)

	if err := http.ListenAndServe(":80", router); err != nil {
		fmt.Print(err)
		return
	}
}
