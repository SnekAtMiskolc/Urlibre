package main

import (
	"context"
	"fmt"

	mongourl "example.com/urlibre/mongo_url"
	"github.com/gin-gonic/gin"
)

func main() {
	mongo_uri := "mongodb://adonis:CoCk1234@localhost:27017"

	controller := mongourl.UrlController{}

	err := controller.Connect(mongo_uri)

	if err != nil {
		panic("Error connecting to mongodb")
	}
	fmt.Println("Connected!")
	defer controller.Client.Disconnect(context.TODO())
	fmt.Println("Done with cleanup!")

	router := gin.Default()

	router.GET("/", index)

	router.Run()
}

func index(ctx *gin.Context) {
	ctx.String(200, "Use curl (or other http client) to send a POST request to /newurl containing the following JSON: 
	{\"url\":\"...\",\"exp\":MAX_IS_120DAYS}")
}
