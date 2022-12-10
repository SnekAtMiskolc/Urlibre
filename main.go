package main

import (
	"context"

	"example.com/urlibre/controllers"
	"example.com/urlibre/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server        *gin.Engine
	urlservice    services.UrlService
	urlcontroller controllers.UrlController
	contex        context.Context
	urlc          *mongo.Collection
	mongoclient   *mongo.Client
	err           error
)

func main() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI()
	// https://piped.video/watch?v=vDIAwtGU9LE&t=3050 <- Continue mangodb X gin gonic guide

}
