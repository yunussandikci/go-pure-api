package common

import (
"github.com/stretchr/testify/assert"
"os"
"testing"
)


func TestUtil_GetPort_DefaultWhenNotProvided(t *testing.T) {
	//given
	os.Unsetenv("PORT")

	//when
	port := GetPort()

	assert.Equal(t, "8080", port)
}

func TestUtil_GetPort_FromEnvironmentVariable(t *testing.T) {
	//given
	os.Setenv("PORT", "101")
	defer os.Unsetenv("PORT")

	//when
	port := GetPort()

	assert.Equal(t, "101", port)
}

func TestUtil_GetConnectionString(t *testing.T) {
	//given
	os.Setenv("MONGO_URI", "foo:bar@tcp(baz)/qux?charset=utf8")
	defer os.Unsetenv("MONGO_URIL")

	//when
	connectionString := GetMongoURI()

	assert.Equal(t, "foo:bar@tcp(baz)/qux?charset=utf8", connectionString)
}
