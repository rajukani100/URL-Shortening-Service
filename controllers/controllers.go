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
		fmt.Fprint(w, "Please provide url", http.StatusBadRequest)
		return
	}

	if _, urlErr := url.ParseRequestURI(modelURL.Url); urlErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide valid url", http.StatusBadRequest)
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
		fmt.Fprint(w, "error while encoding", http.StatusBadRequest)
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
