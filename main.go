package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client *mongo.Client
)

func main() {
	connectDB()
	r := setupRouter()
	r.Run(":" + ginPortYarbDB)
	//	disconnectDB()

}

func mongoUpdateIGPost(name string, ts int) User {
	collection := client.Database("yarb").Collection("users")

	// Update a document
	filter := bson.D{{"name", name}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"date.instagram_post", ts}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//Find a single document
	var result User

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func mongoUpdateIGStories(name string, ts int) User {
	collection := client.Database("yarb").Collection("users")

	// Update a document
	filter := bson.D{{"name", name}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"date.instagram_stories", ts}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//Find a single document
	var result User

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func mongoUpdateIGStoriesByID(id int, ts int) User {
	collection := client.Database("yarb").Collection("users")

	// Update a document
	filter := bson.D{{"social.instagram_id", id}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"date.instagram_stories", ts}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//Find a single document
	var result User

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func Find() []*User {
	// Get a handle for your collection
	collection := client.Database("yarb").Collection("users")

	filter := bson.D{{"setting.disabled", false}}
	findOptions := options.Find()
	findOptions.SetLimit(0)

	var results []*User

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return []*User{}
		//log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	//log.Printf("%v", reflect.TypeOf(results[0]))
	//log.Printf("%v", results[0])

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}

func FindUser(name string) *User {
	// Get a handle for your collection
	collection := client.Database("yarb").Collection("users")

	//	findOptions := options.Find()
	//	findOptions.SetLimit(1)
	filter := bson.D{{"name", name}}

	var result *User

	// Finding multiple documents returns a cursor
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return &User{}
		log.Fatal(err)
	}

	return result
}

func mongoGetUserByIstagramID(id int) *User {
	// Get a handle for your collection
	collection := client.Database("yarb").Collection("users")

	//	findOptions := options.Find()
	//	findOptions.SetLimit(1)
	filter := bson.D{{"social.instagram_id", id}}

	var result *User

	// Finding multiple documents returns a cursor
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return &User{}
		log.Fatal(err)
	}

	return result
}

func mongoGetIGPostTs(name string) int {
	// Get a handle for your collection
	collection := client.Database("yarb").Collection("users")

	//	findOptions := options.Find()
	//	findOptions.SetLimit(1)
	filter := bson.D{{"name", name}}

	var result *User

	// Finding multiple documents returns a cursor
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return 0
		log.Fatal(err)
	}

	return result.Date.InstagramPost
}

func mongoGetIGStoriesTs(name string) int {
	// Get a handle for your collection
	collection := client.Database("yarb").Collection("users")

	//	findOptions := options.Find()
	//	findOptions.SetLimit(1)
	filter := bson.D{{"name", name}}

	var result *User

	// Finding multiple documents returns a cursor
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return 0
		log.Fatal(err)
	}

	return result.Date.InstagramStories
}

func mongoGetIGStoriesTsByID(id int) int {
	// Get a handle for your collection
	collection := client.Database("yarb").Collection("users")

	//	findOptions := options.Find()
	//	findOptions.SetLimit(1)
	filter := bson.D{{"social.instagram_id", id}}

	var result *User

	// Finding multiple documents returns a cursor
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return 0
		log.Fatal(err)
	}
	//println("result:", result)

	return result.Date.InstagramStories
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
