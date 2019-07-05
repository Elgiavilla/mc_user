package main

import (
	"fmt"
	"time"
	"os"

	"github.com/spf13/viper"
	"github.com/elgiavilla/mc_user/middleware"
	_http "github.com/elgiavilla/mc_user/users/http"
	_repo "github.com/elgiavilla/mc_user/users/repository"
	_svc "github.com/elgiavilla/mc_user/users/service"
	"github.com/juju/mgosession"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

func init(){
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`){
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	info := &mgo.DialInfo{
		Addrs:    []string{os.Getenv(`DB_HOST`)},
		Timeout:  60 * time.Second,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	session, err = mgo.Dial(os.Getenv(`DB_HOST`))
	if err != nil {
		fmt.Println(err)
	}
	mPool := mgosession.NewPool(nil, session, 10)
	defer mPool.Close()

	timeoutContext := time.Duration(1) * time.Second
	userRepo := _repo.NewMongoRepo(mPool, os.Getenv(`DB_NAME`))
	userSvc := _svc.NewService(userRepo, timeoutContext)
	e := echo.New()
	mddl := middleware.InitMiddleware()
	e.Use(mddl.CORS)
	_http.NewUserHandler(e, userSvc)
	e.Start(viper.GetString(`server.address`))
}