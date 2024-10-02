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
	router.HandleFunc("PUT /shorten/{shortenCode}", controllers.UpdateShorten)
	router.HandleFunc("DELETE /shorten/{shortenCode}", controllers.DeleteShorten)
	router.HandleFunc("GET /shorten/{shortenCode}/stats", controllers.StatsShorten)

	if err := http.ListenAndServe(":80", router); err != nil {
		fmt.Print(err)
		return
	}
}
