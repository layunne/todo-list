package application

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/layunne/todo-list/backend/controller"
)

func NewWebServer(
	port string,
	usersWebController controller.UsersWebController,
	projectsWebController controller.ProjectsWebController,
	tasksWebController controller.TasksWebController,
) Server {
	return &webServer{usersWebController: usersWebController, projectsWebController: projectsWebController, tasksWebController: tasksWebController, port: port}
}

type webServer struct {
	usersWebController controller.UsersWebController
	projectsWebController controller.ProjectsWebController
	tasksWebController controller.TasksWebController
	port               string
}

func (w *webServer) Start() error {

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())


	e.GET("/users", func(context echo.Context) error { return w.usersWebController.OnGet(context) })
	e.POST("/users", func(context echo.Context) error { return w.usersWebController.OnCreate(context) })
	e.PUT("/users", func(context echo.Context) error { return w.usersWebController.OnUpdate(context) })
	e.DELETE("/users", func(context echo.Context) error { return w.usersWebController.OnDelete(context) })
	e.POST("/users/login", func(context echo.Context) error { return w.usersWebController.OnLogin(context) })

	e.GET("/projects/:id", func(context echo.Context) error { return w.projectsWebController.OnGet(context) })
	e.GET("/projects", func(context echo.Context) error { return w.projectsWebController.OnGetAll(context) })
	e.POST("/projects", func(context echo.Context) error { return w.projectsWebController.OnCreate(context) })
	e.PUT("/projects", func(context echo.Context) error { return w.projectsWebController.OnUpdate(context) })
	e.DELETE("/projects/:id", func(context echo.Context) error { return w.projectsWebController.OnDelete(context) })

	e.GET("/tasks/:id", func(context echo.Context) error { return w.tasksWebController.OnGet(context) })
	e.GET("/tasks/all/:projectId", func(context echo.Context) error { return w.tasksWebController.OnGetAll(context) })
	e.POST("/tasks", func(context echo.Context) error { return w.tasksWebController.OnCreate(context) })
	e.PUT("/tasks", func(context echo.Context) error { return w.tasksWebController.OnUpdate(context) })
	e.PUT("/tasks/status", func(context echo.Context) error { return w.tasksWebController.OnChangeStatus(context) })
	e.DELETE("/tasks/:id", func(context echo.Context) error { return w.tasksWebController.OnDelete(context) })

	return e.Start(fmt.Sprintf(":%s", w.port))
}
