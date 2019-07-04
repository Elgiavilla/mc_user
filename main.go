package main

import (
	"fmt"
	"time"

	"github.com/elgiavilla/mc_user/middleware"

	"github.com/elgiavilla/mc_user/config"
	_http "github.com/elgiavilla/mc_user/users/http"
	_repo "github.com/elgiavilla/mc_user/users/repository"
	_svc "github.com/elgiavilla/mc_user/users/service"
	"github.com/juju/mgosession"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		fmt.Println(err)
	}
	mPool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTIONPOOL)
	defer mPool.Close()

	timeoutContext := time.Duration(1) * time.Second
	userRepo := _repo.NewMongoRepo(mPool, config.MONGODB_DATABASE)
	userSvc := _svc.NewService(userRepo, timeoutContext)
	e := echo.New()
	mddl := middleware.InitMiddleware()
	e.Use(mddl.CORS)
	_http.NewUserHandler(e, userSvc)
	e.Start(":8000")
}
