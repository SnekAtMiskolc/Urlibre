package mongourl

import (
	"context"

	"example.com/urlibre/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlController struct {
	Coll *mongo.Collection
	Ctx  context.Context
}

func (c *UrlController) New(coll *mongo.Collection, ctx context.Context) *UrlController {
	return &UrlController{
		Coll: coll,
		Ctx:  ctx,
	}
}

func (c *UrlController) InsertUrl(url *models.URL) (*models.URL, error) {
	_, err := c.Coll.InsertOne(c.Ctx, url)
	return url, err
}

func (c *UrlController) GetUrl(urlID string) (*models.URL, error) {
	var url *models.URL
	query := bson.D{bson.E{Key: "id", Value: urlID}}
	err := c.Coll.FindOne(c.Ctx, query).Decode(&url)
	return url, err
}
