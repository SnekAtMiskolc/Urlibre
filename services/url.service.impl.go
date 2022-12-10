package services

import (
	"context"

	"example.com/urlibre/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlServiceImpl struct {
	urlCollection *mongo.Collection
	ctx           context.Context
}

func (u *UrlServiceImpl) CreateUrl(url *models.Url) error {
	_, err := u.urlCollection.InsertOne(u.ctx, url)
	return err
}

func (u *UrlServiceImpl) GetUrl(id *string) (*models.Url, error) {
	var url *models.Url

	query := bson.D{bson.E{Key: "id", Value: id}}
	err := u.urlCollection.FindOne(u.ctx, query).Decode(&url)

	return url, err
}
