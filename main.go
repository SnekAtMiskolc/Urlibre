package main

import (
	"context"
	"fmt"

	mongourl "example.com/urlibre/mongo_url"
	//"github.com/gin-gonic/gin"
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
}
