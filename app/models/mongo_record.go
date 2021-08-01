package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoRecord struct {
	Id         primitive.ObjectID `bson:"_id"`
	Key        string             `bson:"key"`
	Value      string             `bson:"value"`
	CreatedAt  time.Time          `bson:"createdAt"`
	TotalCount int                `bson:"totalCount"`
}
