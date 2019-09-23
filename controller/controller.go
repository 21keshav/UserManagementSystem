package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/21keshav/UserManagementSystem/controller/resources"
	"github.com/21keshav/UserManagementSystem/util"
	"github.com/golang/glog"
	"github.com/labstack/echo"
)

type Controller interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	AttachHandlers(lister *echo.Echo)
}

type ControllerImpl struct {
	UserManager resources.UserManager
}

func NewController(userManager resources.UserManager) Controller {
	return &ControllerImpl{
		userManager,
	}
}

func (co *ControllerImpl) AttachHandlers(lister *echo.Echo) {
	lister.GET("/user", co.GetUser)
	lister.POST("/create-user", co.CreateUser)
}

func (co *ControllerImpl) GetUser(c echo.Context) error {
	glog.Info("get-user")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")

	id := c.QueryParam("id")
	user, err := co.UserManager.GetUser(id)
	if err != nil {
		glog.Error("get-user-error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (co *ControllerImpl) CreateUser(c echo.Context) error {

	glog.Info("create-user")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")

	var userDetails util.User
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		glog.Error("read-error", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	err = json.Unmarshal(body, &userDetails)
	if err != nil {
		glog.Error("unmarshal-error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := co.UserManager.CreateUser(userDetails)
	if err != nil {
		glog.Error("create-user-error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println(id)
	return c.JSON(http.StatusCreated, id)
}
