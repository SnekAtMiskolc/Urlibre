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
	// General CTX for most tasks
	ctx = context.Background()

	/*
		MongoDB URI planned to just use an env in the future but since i have to
		develop a few little features so it stays for now!
	*/
	mongo_uri = "mongodb://127.0.0.1:27017"

	mongoconn := options.Client().ApplyURI(mongo_uri)

	// Connecting to mongoDB
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	// Testing if mongo works!
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

	// Attach Url related Services
	router = us.AttachUrlServices(router)
	// Attach Index for a minimal guide for usage!
	router.GET("/", Index)

	router.Run()
}

func Index(ctx *gin.Context) {
	ctx.String(200, "URLibre \n Use curl (or other http client) to send a POST request to /newurl containing the following JSON: \n{url:Hello,exp:30} \n Source Code: https://gitlab.com/GithubSucks/urlibre")
}
