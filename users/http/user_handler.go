package http

import (
	"net/http"

	"github.com/elgiavilla/mc_user/models"
	"github.com/elgiavilla/mc_user/users"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ReponseError struct {
	Message string `json:"message"`
}

type HttpMongo struct {
	MongoService users.Service
}

func NewUserHandler(e *echo.Echo, UserService users.Service) {
	handler := &HttpMongo{
		MongoService: UserService,
	}
	e.GET("/users", handler.FindAll)
	e.GET("/user/:id", handler.Find)
	e.POST("/user", handler.Store)
	e.DELETE("/user/:id", handler.Delete)
}

func (n *HttpMongo) FindAll(c echo.Context) error {
	list, err := n.MongoService.FindAll()
	if err != nil {
		c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (n *HttpMongo) Store(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	d, err := n.MongoService.Store(&user)
	return c.JSON(http.StatusOK, d)
}

func (n *HttpMongo) Find(c echo.Context) error {
	id := c.Param("id")
	idP := models.StringToID(id)
	list, err := n.MongoService.Find(idP)
	if err != nil {
		c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (n *HttpMongo) Delete(c echo.Context) error {
	id := c.Param("id")
	idP := models.StringToID(id)
	_ = n.MongoService.Delete(idP)
	return c.JSON(http.StatusOK, "OK")
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
