package main

import "github.com/yunussandikci/go-pure-api/app"

// @title Go Pure API
// @version 1.0
// @description A REST API that allows you to get records from mongo database and read/write them to in-memory database!
// @schemes https
// @host go-pure-api.herokuapp.com
// @BasePath /api/v1
func main() {
	application := app.NewApp()
	defer application.Stop()
	application.Run()
}
