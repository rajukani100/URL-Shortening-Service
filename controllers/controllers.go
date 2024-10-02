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

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateShorten(w http.ResponseWriter, r *http.Request) {
	var modelURL models.URL

	//retrieve url
	if err := json.NewDecoder(r.Body).Decode(&modelURL); err != nil {
		fmt.Fprint(w, "Please provide url", http.StatusBadRequest)
		return
	}

	if _, urlErr := url.ParseRequestURI(modelURL.Url); urlErr != nil {
		fmt.Fprint(w, "Please provide valid url", http.StatusBadRequest)
		return
	}

	hash := md5.Sum([]byte(modelURL.Url))
	shorten := hex.EncodeToString(hash[:])[:6]

	url_info := &models.URL_INFO{
		Id:        primitive.NewObjectID(),
		Url:       modelURL.Url,
		Shorten:   shorten,
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
		fmt.Fprint(w, "error while encoding", http.StatusBadRequest)
		return
	}
}
