package in_memory

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunussandikci/go-pure-api/app/mocks"
	"testing"
)

func TestRecordsRepository_Create(t *testing.T) {
	//given
	inMemoryDatabase := new(mocks.InMemoryDatabase)
	repository := NewInMemoryRecordsRepository(inMemoryDatabase)
	inMemoryDatabase.On("GetSync", "100").Return("")
	inMemoryDatabase.On("CreateSync", "100", "200")
	//when
	createErr := repository.Create("100", "200")

	//then
	assert.NoError(t, createErr)
}

func TestRecordsRepository_Create_KeyAlreadyExistError(t *testing.T) {
	//given
	inMemoryDatabase := new(mocks.InMemoryDatabase)
	repository := NewInMemoryRecordsRepository(inMemoryDatabase)
	inMemoryDatabase.On("GetSync", "100").Return("200")
	inMemoryDatabase.On("CreateSync", "100", "300")
	//when
	createErr := repository.Create("100", "300")

	//then
	assert.Error(t, createErr)
	assert.Contains(t, createErr.Error(), KeyAlreadyExistError)
}

func TestRecordsRepository_Get(t *testing.T) {
	//given
	inMemoryDatabase := new(mocks.InMemoryDatabase)
	repository := NewInMemoryRecordsRepository(inMemoryDatabase)
	inMemoryDatabase.On("GetSync", "100").Return("200")
	//when
	value, createErr := repository.Get("100")

	//then
	assert.Equal(t, value, "200")
	assert.NoError(t, createErr)
}

func TestRecordsRepository_Get_KeyNotExistError(t *testing.T) {
	//given
	inMemoryDatabase := new(mocks.InMemoryDatabase)
	repository := NewInMemoryRecordsRepository(inMemoryDatabase)
	inMemoryDatabase.On("GetSync", "100").Return("")

	//when
	value, createErr := repository.Get("100")

	//then
	assert.Error(t, createErr)
	assert.Contains(t, createErr.Error(), KeyNotExistError)
	assert.Empty(t, value)
}
