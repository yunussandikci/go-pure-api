package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yunussandikci/go-pure-api/app/common"
	v1 "github.com/yunussandikci/go-pure-api/app/controllers/v1"
	"github.com/yunussandikci/go-pure-api/app/database"
	_ "github.com/yunussandikci/go-pure-api/app/docs"
	"github.com/yunussandikci/go-pure-api/app/server"
)

type App struct {
	applicationServer  *server.ApplicationServer
	mongoDatabase      database.MongoDatabase
	inMemoryDatabase   database.InMemoryDatabase
	mongoController    v1.MongoController
	inMemoryController v1.InMemoryController
}

//initDependencies Initializes necessary dependencies of the application.
func (a *App) initDependencies() {
	a.applicationServer = server.NewApplicationServer()
	a.mongoDatabase = database.NewMongoDatabase()
	a.inMemoryDatabase = database.NewInMemoryDatabase()
	a.mongoController = v1.NewMongoController(a.mongoDatabase)
	a.inMemoryController = v1.NewInMemoryController(a.inMemoryDatabase)
}

//initDependencies Initializes route endpoints of the application.
func (a *App) initRoutes() {
	a.applicationServer.HandleFunctions("/api/v1/mongo", server.HandlerFunctions{
		Post: a.mongoController.GetRecords,
	})
	a.applicationServer.HandleFunctions("/api/v1/in-memory", server.HandlerFunctions{
		Get:  a.inMemoryController.GetRecords,
		Post: a.inMemoryController.Create,
	})
	a.applicationServer.HandleFunc("/docs/", httpSwagger.WrapHandler)
}

//Run Runs the application.
func (a *App) Run() {
	port := common.GetPort()
	common.Logger.Infof("🚀 Application running at port:%s", port)
	err := a.applicationServer.Run(port)
	if err != nil {
		panic(err)
	}
}

//Stop Stops the application.
func (a *App) Stop() {
	a.mongoDatabase.Disconnect()
}

func NewApp() *App {
	app := &App{}
	app.initDependencies()
	app.initRoutes()
	return app
}
