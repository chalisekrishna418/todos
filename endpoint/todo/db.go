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
