package database

import (
	"github.com/stretchr/testify/assert"
	testing2 "github.com/yunussandikci/go-pure-api/app/common/testing"
	"testing"
)

func TestMongoDatabase(t *testing.T) {
	//given
	var mongoDatabase MongoDatabase

	//when
	testing2.RunWithMongoTestContainer(func() {
		mongoDatabase = NewMongoDatabase()
	})

	//then
	assert.NotNil(t, mongoDatabase)
	assert.NotNil(t, mongoDatabase.GetCollection("database"))
}