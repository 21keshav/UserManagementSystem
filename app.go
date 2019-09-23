package main

import (
	"context"
	"fmt"
	"time"

	"github.com/21keshav/UserManagementSystem/config"
	"github.com/21keshav/UserManagementSystem/controller"
	"github.com/21keshav/UserManagementSystem/controller/resources"
	"github.com/21keshav/UserManagementSystem/util"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	glog.Info("main")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	e := echo.New()
	var conf config.Config
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		glog.Error("Error reading config file")
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	mongoURL := fmt.Sprintf("%s:%s", conf.Database.Server, conf.Database.Port)
	mongoClient := util.NewMongoClient(context.TODO(), mongoURL)
	context, _ := context.WithTimeout(context.Background(), 60*time.Second)
	userManager := resources.NewUserManager("sbap", "users", mongoClient, context)
	controller := controller.NewController(userManager)
	controller.AttachHandlers(e)
	e.Start(":1234")
}
