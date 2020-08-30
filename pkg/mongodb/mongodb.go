package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBConnection interface
type DBConnection interface {
	Connect() (*mongo.Database, error)
}

//SessionManager mongo session manager
type SessionManager struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

//Connect provides mongodb client
func (ms *SessionManager) Connect() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+ms.Host+":"+ms.Port))
	if err != nil {
		return nil, err
	}
	db := client.Database(ms.Database)
	return db, nil
}
