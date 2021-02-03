package main

import (
	"fmt"
	mongo "github.com/zhaopengme/go-mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	options := mongo.MgoOptions{
		DbHost:   "mongodb://47.252.86.134:27017",
		DbName:   "hotflow",
		Username: "pro",
		Password: "C*fNQNBeY9KAHUp7afYw",
	}
	mongo.MustConnect(options)
}

func main() {
	value := mongo.Instance.InsertOne("agent", bson.M{
		"username": "gufeng",
	})
	fmt.Printf("%v", value)
}
