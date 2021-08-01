package mongo

import (
	"context"
	"github.com/yunussandikci/go-pure-api/app/database"
	"github.com/yunussandikci/go-pure-api/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoRecordsRepository interface {
	GetRecordsWith(minCount int, maxCount int, startDate time.Time, endDate time.Time) ([]models.MongoRecord, error)
}

type mongoRecordsRepository struct {
	mongoDatabase database.MongoDatabase
}

func NewMongoRecordsRepository(database database.MongoDatabase) MongoRecordsRepository {
	return &mongoRecordsRepository{
		mongoDatabase: database,
	}
}

//GetRecordsWith Returns Records from mongodb with the filters provided as minCount, maxCount, startDate and endDate or
//returns an error
func (r *mongoRecordsRepository) GetRecordsWith(minCount int, maxCount int, startDate time.Time, endDate time.Time) ([]models.MongoRecord, error) {
	projection := bson.M{"$project": bson.M{
		"id": 1, "key": 1, "value": 1, "createdAt": 1, "totalCount": bson.M{
			"$sum": "$counts",
		},
	}}
	filter := bson.M{"$match": bson.M{"$and": []bson.M{
		{"createdAt": bson.M{"$gte": startDate}},
		{"createdAt": bson.M{"$lte": endDate}},
		{"totalCount": bson.M{"$gte": minCount}},
		{"totalCount": bson.M{"$lte": maxCount}},
	}}}
	find, findErr := r.getCollection().Aggregate(context.Background(), []bson.M{projection, filter})
	if findErr != nil {
		return nil, findErr
	}
	defer func(find *mongo.Cursor) {
		_ = find.Close(context.Background())
	}(find)
	var mongoRecords []models.MongoRecord
	findAllErr := find.All(context.Background(), &mongoRecords)
	if findAllErr != nil {
		return nil, findAllErr
	}
	return mongoRecords, nil
}

//getCollection Returns a records collection from the database
func (r *mongoRecordsRepository) getCollection() *mongo.Collection {
	return r.mongoDatabase.GetCollection("records")
}