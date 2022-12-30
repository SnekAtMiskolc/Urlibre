package models

type URL struct {
	ID      string `json:"id" bson:"id"`
	Url     string `json:"url" bson:"url"`
	Expires int64  `json:"exp" bson:"exp"`
}
