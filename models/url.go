package models

import (
	"time"

	"github.com/lithammer/shortuuid/v4"
)

// Used to represent a url!
type URL struct {
	ID      string `json:"id" bson:"id"`
	Url     string `json:"url" bson:"url"`
	Expires int64  `json:"exp" bson:"exp"`
}

// Used to represent the creation of a url!
type NewURL struct {
	Url     string `json:"url"`
	Expires int    `json:"exp"`
}

func (nu *NewURL) IntoURL() *URL {
	return &URL{
		ID:      shortuuid.New(),
		Url:     nu.Url,
		Expires: addDays(nu.Expires),
	}
}

func addDays(days int) int64 {
	return time.Now().AddDate(0, 0, days).Unix()
}
