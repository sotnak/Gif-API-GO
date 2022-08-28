package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() *mongo.Client {

	mongo_url := os.Getenv("MONGO_URL")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))

	if err != nil {
		panic(err)
	}

	return client
}

func disconnect(client *mongo.Client) {
	client.Disconnect(context.TODO())
}

func getTagsCount(query string) int64 {
	client := connect()
	defer disconnect(client)

	filter := bson.D{}

	if len(query) > 0 {
		filter = append(filter, bson.E{Key: "name", Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: query, Options: "i"}},
		}},
		)
	}

	count, err := client.Database("nsfw").Collection("tags").CountDocuments(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return count
}

func getTags(query string, limit int64, skip int64) []Tag {
	client := connect()
	defer disconnect(client)

	opts := options.Find().
		SetSort(bson.D{{Key: "count", Value: -1}}).
		SetLimit(limit).
		SetSkip(skip)

	filter := bson.D{}

	if len(query) > 0 {
		filter = append(filter, bson.E{Key: "name", Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: query, Options: "i"}},
		}},
		)
	}

	cursor, err := client.Database("nsfw").Collection("tags").Find(context.TODO(), filter, opts)

	if err != nil {
		panic(err)
	}

	var results []Tag
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if results == nil {
		return []Tag{}
	}

	return results
}

func getGifsByTag(tag string, limit int64, skip int64) []Gif {
	client := connect()
	defer disconnect(client)

	opts := options.Find().
		SetLimit(limit).
		SetSkip(skip)

	filter := bson.D{}

	if len(tag) > 0 {
		filter = append(filter, bson.E{Key: "tags", Value: tag})
	}

	cursor, err := client.Database("nsfw").Collection("gifs").Find(context.TODO(), filter, opts)

	if err != nil {
		panic(err)
	}

	var results []Gif
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if results == nil {
		return []Gif{}
	}

	return results
}

func getRandomGifs(limit int64) []Gif {
	client := connect()
	defer disconnect(client)

	pipeline := []bson.D{{{Key: "$sample", Value: bson.D{{Key: "size", Value: limit}}}}}

	cursor, err := client.Database("nsfw").Collection("gifs").Aggregate(context.TODO(), pipeline)

	if err != nil {
		panic(err)
	}

	var results []Gif
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if results == nil {
		return []Gif{}
	}

	return results
}