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

	//define models
	url_info := &models.URL_INFO{
		Url:       modelURL.Url,
		ShortCode: shorten,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	var stats models.Stats
	stats.Id = primitive.NewObjectID()
	stats.Url_info = *url_info

	collection := database.GetStatsCollection()
	_, err := collection.InsertOne(context.TODO(), stats)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	EncodeErr := json.NewEncoder(w).Encode(ShortenResponse(&stats))
	if EncodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}
}

func RetrieveShorten(w http.ResponseWriter, r *http.Request) {
	shortenCode := r.PathValue("shortenCode")

	//define model
	var stats models.Stats
	stats.Id = primitive.NewObjectID()

	collection := database.GetStatsCollection()
	err := collection.FindOne(context.TODO(), bson.M{"url_info.shorten": shortenCode}).Decode(&stats)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	// increase view
	stats.Views = stats.Views + 1
	_, updateErr := collection.UpdateOne(context.Background(), bson.M{"url_info.shorten": shortenCode}, bson.M{"$set": bson.M{"views": stats.Views}})
	if updateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while updating.")
		return
	}

	//success
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ShortenResponse(&stats)); err != nil {
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
	collection := database.GetStatsCollection()

	// Find and update the document
	filter := bson.M{"url_info.shorten": shortenCode}
	update := bson.M{"$set": bson.M{"url_info.url": modelURL.Url, "url_info.updated_at": time.Now()}}

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
	var updatedStats models.Stats
	err := collection.FindOne(context.TODO(), filter).Decode(&updatedStats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error fetching updated document.")
		return
	}

	// Return the updated URL info
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(ShortenResponse(&updatedStats)); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while encoding response.")
		return
	}
}

func DeleteShorten(w http.ResponseWriter, r *http.Request) {
	shortenCode := r.PathValue("shortenCode")

	//get Mongo collection
	collection := database.GetStatsCollection()

	//delete document
	_, DeleteErr := collection.DeleteOne(
		context.TODO(),
		bson.M{"url_info.shorten": shortenCode})
	if DeleteErr != nil {
		if DeleteErr == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Not found")
			return
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func StatsShorten(w http.ResponseWriter, r *http.Request) {
	shortenCode := r.PathValue("shortenCode")

	collection := database.GetStatsCollection()
	var FoundStats models.Stats
	FindErr := collection.FindOne(context.TODO(), bson.M{"url_info.shorten": shortenCode}).Decode(&FoundStats)
	if FindErr != nil {
		if FindErr == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "user not found.")
		}
		return

	}

	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(FoundStats)
	if encodeErr != nil {
		fmt.Fprint(w, "Error while encoding.")
		return
	}

}
