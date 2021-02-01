package main

import (
	"github.com/layunne/todo-list/backend/application"
	"github.com/layunne/todo-list/backend/config"
	"github.com/layunne/todo-list/backend/controller"
	"github.com/layunne/todo-list/backend/infrastructure/mongo"
	"github.com/layunne/todo-list/backend/repositories"
	"github.com/layunne/todo-list/backend/services"
	"log"
)

func main() {

	env := config.NewEnv()

	mongoDb := mongo.NewMongoClient(env.MongoHost(), env.MongoDatabase())

	usersRepository := repositories.NewUsersRepository(mongoDb)
	projectsRepository := repositories.NewProjectsRepository(mongoDb)
	tasksRepository := repositories.NewTasksRepository(mongoDb)

	authService := services.NewAuthService(env.AuthSecret())
	encryptionService := services.NewEncryptionService()
	usersService := services.NewUsersService(usersRepository, authService, encryptionService)
	projectsService := services.NewProjectsService(projectsRepository, tasksRepository)
	tasksService := services.NewTasksService(tasksRepository, projectsRepository)

	usersWebController := controller.NewUsersWebController(usersService)
	projectsWebController := controller.NewProjectsWebController(projectsService, authService)
	tasksWebController := controller.NewTasksWebController(tasksService, authService)

	webServer := application.NewWebServer(env.WebServerPort(), usersWebController, projectsWebController, tasksWebController)

	servers := []application.Server{webServer}

	exit := make(chan error)
	for _, server := range servers {
		go func(server application.Server) {
			exit <- server.Start()
		}(server)
	}

	log.Fatalln(<-exit)
}
