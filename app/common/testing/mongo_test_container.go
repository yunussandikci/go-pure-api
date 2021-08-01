package testing

import (
	"fmt"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
)

func RunWithMongoTestContainer(f func()) {
	mongoPort := "27017/tcp"
	mongoUsername := "user"
	mongoPassword := "pass"
	mongoDatabase := "database"

	container := NewTestContainer("mongo:latest", []string{mongoPort}, map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": mongoUsername,
		"MONGO_INITDB_ROOT_PASSWORD": mongoPassword,
		"MONGO_INITDB_DATABASE":      mongoDatabase,
	}, wait.ForLog("Listening on 0.0.0.0"))

	hostPort := container.GetHostPort(mongoPort)
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s", mongoUsername, mongoPassword, hostPort)

	_ = os.Setenv("MONGO_URI", mongoUri)
	defer os.Unsetenv("MONGO_URI")
	defer container.Terminate()
	f()
}
