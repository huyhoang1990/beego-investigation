package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/huyhoang1990/beego-investigation/conf"
	"github.com/huyhoang1990/beego-investigation/controllers"
	"github.com/huyhoang1990/beego-investigation/infras"
	"github.com/huyhoang1990/beego-investigation/repo/mysql"
	"github.com/huyhoang1990/beego-investigation/service"
	"github.com/sirupsen/logrus"
)

var LogRusss = logrus.New()

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*30, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	LogRusss.Out = os.Stdout

	go func() {

		db, err := infras.NewMysqlSession()
		if err != nil {
			panic(err)
		}
		userRepo := mysql.NewUserRepo(db)
		passwordService := service.NewPasswordService()

		authService := service.NewAuthService(userRepo, passwordService)
		baseController := &controllers.BaseController{
			UserRepo: userRepo,
		}
		homeController := controllers.NewHomeController(baseController)
		authController := controllers.NewAuthenticationController(*baseController, authService)

		web.Router(conf.HOME_PATH, homeController, "get:Home")
		web.Router(conf.LOGIN_PATH, authController, "post:Login;get:Login")

		web.BConfig.EnableErrorsShow = true
		web.BConfig.WebConfig.ViewsPath = "static"
		web.BConfig.Listen.ServerTimeOut = 5
		web.BConfig.WebConfig.StaticDir["/static"] = "static"

		web.Run()

	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	web.BeeApp.Server.Shutdown(ctx)
	fmt.Println("shutting down")
	log.Println("shutting down")
	LogRusss.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	os.Exit(0)
}
