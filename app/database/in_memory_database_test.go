package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSync(t *testing.T) {
	//given
	inMemoryDatabase := NewInMemoryDatabase()

	//when
	inMemoryDatabase.CreateSync("100","101")

	//then
	value := inMemoryDatabase.GetSync("100")
	assert.Equal(t, "101", value)
}

func TestGetSync(t *testing.T) {
	//given
	inMemoryDatabase := NewInMemoryDatabase()

	//when
	inMemoryDatabase.CreateSync("100","101")

	//then
	value := inMemoryDatabase.GetSync("100")
	assert.Equal(t, "101", value)
}