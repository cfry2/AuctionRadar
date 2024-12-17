package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuctionItem struct {
	Name     string `json:"name" binding:"required"`
	FeedLink string `json:"feedLink" binding:"required"`
}

type AuctionItems []AuctionItem

var dbclient *mongo.Client
var err error

func getItemsHandler(c *gin.Context) {
	var results AuctionItems
	cursor, err := dbclient.Database("AuctionRadarDB").Collection("Auction_Items").Find(context.TODO(), gin.H{}, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "NOT OK"})
		panic(err)
	}
	cursor.All(context.TODO(), &results)
	c.JSON(http.StatusOK, gin.H{"message": results})
}

func addItemsHandler(c *gin.Context) {
	var auctionItem AuctionItem
	if c.ShouldBind(&auctionItem) == nil {
		dbclient.Database("AuctionRadarDB").Collection("Auction_Items").InsertOne(context.TODO(), auctionItem)
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}

}

func deleteItemsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from BFF"})
}

// main is the entry point for the application.
//
// It sets up a gin router that listens on port 4000 and
// responds to GET requests to /items with a JSON response.
func main() {
	r := gin.Default()
	dbclient, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Running BFF....")
	r.GET("/items", getItemsHandler)
	r.PUT("/items", addItemsHandler)
	r.DELETE("/items", deleteItemsHandler)
	r.Run(":4000")
}
