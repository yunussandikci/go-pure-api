package common

import (
	"fmt"
	"os"
	"strings"
)

//GetPort Resolves port that application needs to be start from environment variables.
func GetPort() string {
	const DefaultPort = "8080"
	port := strings.TrimSpace(os.Getenv("PORT"))
	if len(port) == 0 {
		port = DefaultPort
	}
	return fmt.Sprintf("%s", port)
}

//GetMongoURI Resolves Mongo URI for the application database from environment variables.
func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}
