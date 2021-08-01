package database

import (
	"context"
	"github.com/yunussandikci/go-pure-api/app/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type MongoDatabase interface {
	GetCollection(collection string) *mongo.Collection
	Disconnect()
}

type mongoDatabase struct {
	client *mongo.Client
	connectionString *connstring.ConnString
}

func (m mongoDatabase) GetCollection(collection string) *mongo.Collection {
	return m.client.Database(m.connectionString.Database).Collection(collection)
}

func NewMongoDatabase() MongoDatabase {
	connectionString, parseErr := connstring.Parse(common.GetMongoURI())
	if parseErr != nil {
		panic(parseErr)
	}

	client, connectErr := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString.Original))
	if connectErr != nil {
		panic(connectErr)
	}
	return &mongoDatabase{client: client, connectionString: &connectionString}
}

func (m mongoDatabase) Disconnect()  {
	_ = m.client.Disconnect(context.Background())
}
