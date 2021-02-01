package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/layunne/todo-list/backend/models"
	"github.com/layunne/todo-list/backend/services"
)

type TasksWebController interface {
	OnGet(echo echo.Context) error
	OnGetAll(echo echo.Context) error
	OnCreate(echo echo.Context) error
	OnUpdate(echo echo.Context) error
	OnChangeStatus(echo echo.Context) error
	OnDelete(echo echo.Context) error
}

func NewTasksWebController(tasksService services.TasksService, authService services.AuthService) TasksWebController {
	return &tasksWebController{tasksService: tasksService, authService: authService}
}

type tasksWebController struct {
	tasksService services.TasksService
	authService services.AuthService
}

func (w *tasksWebController) OnGet(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	taskId := c.Param("id")

	project := w.tasksService.Get(id, taskId)

	if project == nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, project)
}

func (w *tasksWebController) OnGetAll(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	projectId := c.Param("projectId")

	tasks := w.tasksService.GetAll(id, projectId)

	return c.JSON(http.StatusOK, tasks)
}

func (w *tasksWebController) OnCreate(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	createTask := &models.CreateTask{}

	if err := c.Bind(createTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
			"error": err.Error(),
		})
	}

	userResp, err := w.tasksService.Create(id, createTask)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *tasksWebController) OnUpdate(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	updateTask := &models.UpdateTask{}

	if err := c.Bind(updateTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	userResp, err := w.tasksService.Update(id, updateTask)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *tasksWebController) OnChangeStatus(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	updateTask := &models.UpdateTask{}

	if err := c.Bind(updateTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	userResp, err := w.tasksService.ChangeStatus(id, updateTask)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *tasksWebController) OnDelete(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	taskId := c.Param("id")

	w.tasksService.Delete(id, taskId)

	return c.NoContent(http.StatusNoContent)
}
