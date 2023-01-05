package mongourl

import (
	"context"
	"time"

	"example.com/urlibre/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// Crontab
	"github.com/robfig/cron"
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

func (c *UrlController) SetupCron() *cron.Cron {
	cr := cron.New()
	cr.AddFunc("0 0 * * *", func() {
		println("RUNNING")
		filter := bson.M{"exp": bson.M{"$lt": time.Now().Unix()}}

		_, err := c.Coll.DeleteMany(c.Ctx, filter)
		if err != nil {
			return
		}
	})
	return cr
}
