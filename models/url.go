package models

type Url struct {
	ID  string `json:"id" bson:"url_id" form:"id"`
	url string `json:"url" bson:"url_url" form:"url"`
	exp string `json:"exp" bson:"url_exp" form:"exp"`
}
