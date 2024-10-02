package controllers

import (
	"UrlShorteningService/database"
	"UrlShorteningService/models"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateShorten(w http.ResponseWriter, r *http.Request) {
	var modelURL models.URL

	//retrieve url
	if err := json.NewDecoder(r.Body).Decode(&modelURL); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide url")
		return
	}

	if _, urlErr := url.ParseRequestURI(modelURL.Url); urlErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide valid url")
		return
	}

	hash := md5.Sum([]byte(modelURL.Url))
	shorten := hex.EncodeToString(hash[:])[:6]

	url_info := &models.URL_INFO{
		Id:        primitive.NewObjectID(),
		Url:       modelURL.Url,
		ShortCode: shorten,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Views:     0,
	}

	collection := database.GetUrlsCollection()
	_, err := collection.InsertOne(context.TODO(), url_info)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	EncodeErr := json.NewEncoder(w).Encode(url_info)
	if EncodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}
}

func RetrieveShorten(w http.ResponseWriter, r *http.Request) {
	shortenCode := r.PathValue("shortenCode")

	//define model
	var url_info models.URL_INFO

	collection := database.GetUrlsCollection()
	err := collection.FindOne(context.TODO(), bson.D{{Key: "shorten", Value: shortenCode}}).Decode(&url_info)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	//success
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(url_info); err != nil {
		return
	}
}

func UpdateShorten(w http.ResponseWriter, r *http.Request) {
	shortenCode := r.PathValue("shortenCode")

	// Defined model
	var modelURL models.URL

	// Decode request body into modelURL
	if err := json.NewDecoder(r.Body).Decode(&modelURL); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while decoding request body")
		return
	}

	// Validate URL
	if _, urlErr := url.ParseRequestURI(modelURL.Url); urlErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide a valid URL")
		return
	}

	// Get MongoDB collection
	collection := database.GetUrlsCollection()

	// Find and update the document
	filter := bson.D{{"shorten", shortenCode}}
	update := bson.M{"$set": bson.M{"url": modelURL.Url, "updated_at": time.Now()}}

	result, updateErr := collection.UpdateOne(context.TODO(), filter, update)
	if updateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while updating.")
		return
	}

	// If no document was found
	if result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Not found")
		return
	}

	// Fetch the updated document
	var updatedURL models.URL_INFO
	err := collection.FindOne(context.TODO(), filter).Decode(&updatedURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error fetching updated document.")
		return
	}

	// Return the updated URL info
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(updatedURL); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while encoding response.")
		return
	}
}
