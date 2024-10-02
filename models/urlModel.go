package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	Url string
}

type URL_INFO struct {
	Url       string    `bson:"url" json:"url"`
	ShortCode string    `bson:"shorten" json:"shorten"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type Stats struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Views    int                `bson:"views" json:"views"`
	Url_info URL_INFO
}

type ShortenResponse struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Url       string             `bson:"url" json:"url"`
	ShortCode string             `bson:"shorten" json:"shorten"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
