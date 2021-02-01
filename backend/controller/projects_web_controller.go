package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/layunne/todo-list/backend/models"
	"github.com/layunne/todo-list/backend/services"
)

type ProjectsWebController interface {
	OnGet(echo echo.Context) error
	OnGetAll(echo echo.Context) error
	OnCreate(echo echo.Context) error
	OnUpdate(echo echo.Context) error
	OnDelete(echo echo.Context) error
}

func NewProjectsWebController(projectsService services.ProjectsService, authService services.AuthService) ProjectsWebController {
	return &projectsWebController{projectsService: projectsService, authService: authService}
}

type projectsWebController struct {
	projectsService services.ProjectsService
	authService services.AuthService
}

func (w *projectsWebController) OnGet(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	projectId := c.Param("id")

	project := w.projectsService.Get(id, projectId)

	if project == nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, project)
}

func (w *projectsWebController) OnGetAll(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	projects := w.projectsService.GetAll(id)

	return c.JSON(http.StatusOK, projects)
}

func (w *projectsWebController) OnCreate(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	createProject := &models.CreateProject{}

	if err := c.Bind(createProject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	createProject.UserId = id

	userResp, err := w.projectsService.Create(createProject)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *projectsWebController) OnUpdate(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	updateProject := &models.UpdateProject{}

	if err := c.Bind(updateProject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	updateProject.UserId = id

	userResp, err := w.projectsService.Update(updateProject)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *projectsWebController) OnDelete(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	id, err := w.authService.GetId(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	projectId := c.Param("id")

	err = w.projectsService.Delete(id, projectId)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.NoContent(http.StatusNoContent)
}
