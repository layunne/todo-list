package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/layunne/todo-list/backend/models"
	"github.com/layunne/todo-list/backend/services"
)

type UsersWebController interface {
	OnGet(echo echo.Context) error
	OnCreate(echo echo.Context) error
	OnUpdate(echo echo.Context) error
	OnDelete(echo echo.Context) error
	OnLogin(echo echo.Context) error
}

func NewUsersWebController(usersService services.UsersService) UsersWebController {
	return &usersWebController{usersService: usersService}
}

type usersWebController struct {
	usersService services.UsersService
}

func (w *usersWebController) OnGet(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	user, err := w.usersService.Get(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	if user == nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, user)
}

func (w *usersWebController) OnCreate(c echo.Context) error {

	createUser := &models.CreateUser{}

	if err := c.Bind(createUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	userResp, err := w.usersService.Create(createUser)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *usersWebController) OnUpdate(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")
	
	user := &models.UpdateUser{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	userResp, err := w.usersService.Update(authorization, user)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *usersWebController) OnDelete(c echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	err := w.usersService.Delete(authorization)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.NoContent(http.StatusOK)
}

func (w *usersWebController) OnLogin(c echo.Context) error {

	userLogin := &models.UserLogin{}

	if err := c.Bind(userLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	userResp, err := w.usersService.Login(userLogin)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}
