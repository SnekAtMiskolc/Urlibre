package mongourl

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UrlController struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func (c *UrlController) Connect(uri string) error {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err == nil {
		c.Client = client
		db := c.Client.Database("URL")
		c.Db = db
	}
	return err
}

func (c *UrlController) InsertUrl() {
	c.Db.Collection("URL").InsertOne(context.TODO())
}

func (c *UrlController) CreateCollection() {
	c.Db.CreateCollection(context.Background(), "URLS")
}
