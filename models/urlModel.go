package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	Url string
}

type URL_INFO struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Url       string             `bson:"url" json:"url"`
	Shorten   string             `bson:"shorten" json:"shorten"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Views     int                `bson:"views" json:"views"`
}
