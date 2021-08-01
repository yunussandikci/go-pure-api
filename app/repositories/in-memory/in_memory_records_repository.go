package in_memory

import (
	"errors"
	"github.com/yunussandikci/go-pure-api/app/database"
)

const KeyAlreadyExistError = "Key already exist."
const KeyNotExistError = "Key not exist."

type InMemoryRecordsRepository interface {
	Get(key string) (string, error)
	Create(key, value string) error
}

type inMemoryRecordsRepository struct {
	database database.InMemoryDatabase
}

func NewInMemoryRecordsRepository(database database.InMemoryDatabase) InMemoryRecordsRepository {
	return &inMemoryRecordsRepository{
		database: database,
	}
}

//Get Returns value of a provided key from database or throws an error if not exist.
func (r *inMemoryRecordsRepository) Get(key string) (string, error) {
	record := r.database.GetSync(key)
	if len(record) == 0 {
		return "", errors.New(KeyNotExistError)
	}
	return record, nil
}

//Create Returns creates a key-value pair in the database or throws an error if already exist.
func (r *inMemoryRecordsRepository) Create(key, value string) error {
	record := r.database.GetSync(key)
	if len(record) != 0 {
		return errors.New(KeyAlreadyExistError)
	}
	r.database.CreateSync(key, value)
	return nil
}
