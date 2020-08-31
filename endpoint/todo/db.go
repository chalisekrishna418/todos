package todo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoConnection Connection interface
type MongoConnection interface {
	Connect() (*mongo.Database, error)
}

//MongoDBManager db manager
type MongoDBManager struct {
	DBMgr MongoConnection
}

//Insert inserts a data to mongodb database
func (mm *MongoDBManager) Insert(collection string, data Todo) error {
	client, err := mm.DBMgr.Connect()
	if err != nil {
		log.Fatalf("Error Connecting mongodb. Error: %v", err)
		return err
	}
	c := client.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = c.InsertOne(ctx, bson.M{"todo": data})
	if err != nil {
		log.Fatalf("Error inserting document %v Error: %v", data, err)
		return err
	}
	return nil
}

//List lists data available in a mongodb collection
func (mm *MongoDBManager) List(collection string) ([]interface{}, error) {
	client, err := mm.DBMgr.Connect()
	if err != nil {
		return nil, err
	}
	c := client.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cur, err := c.Find(ctx, bson.D{})
	if err != nil {
		defer cur.Close(ctx)
		return nil, err
	}
	var todos []interface{}
	for cur.Next(ctx) {
		var result bson.M
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		todos = append(todos, result["todo"])
	}
	return todos, nil
}
