package main

import (
	"context"
	"fmt"
	"log"

	mongourl "example.com/urlibre/mongo_url"
	urlservice "example.com/urlibre/url_service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	uc          mongourl.UrlController
	us          urlservice.UrlService
	ctx         context.Context
	urlc        *mongo.Collection
	mongoclient *mongo.Client
	err         error
	mongo_uri   string
)

func init() {
	ctx = context.TODO()

	mongo_uri = "mongodb://127.0.0.1:27017"

	mongoconn := options.Client().ApplyURI(mongo_uri)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	urlc = mongoclient.Database("URL").Collection("urls")
	uc = *uc.New(urlc, ctx)
	us = us.New(uc)
	server = gin.Default()
}

func main() {

	router := gin.Default()

	router = us.AttachUrlServices(router)
	router.GET("/", index)

	router.Run()
}

func index(ctx *gin.Context) {
	ctx.String(200, "Use curl (or other http client) to send a POST request to /newurl containing the following JSON: \n{url:...,exp:MAX_IS_120DAYS}")
}
