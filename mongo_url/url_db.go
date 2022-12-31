package mongourl

import (
	"context"

	"example.com/urlibre/models"
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

/*
func (c *UrlController) InsertUrl() {
	c.Db.Collection("URL").InsertOne(context.TODO(), interface{})
}
*/

func (c *UrlController) CreateUrl(urlData models.URL) {
}
