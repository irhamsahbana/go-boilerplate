package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-boilerplate/bootstrap"
	_userHttp "go-boilerplate/domain/user/delivery/http"
	_userRepo "go-boilerplate/domain/user/repository/mongo"
	_userUsecase "go-boilerplate/domain/user/usecase"
)

func main() {
	if !bootstrap.App.Config.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())

	timeoutContext := time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second
	mongoDatabase := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongodb.name"))

	userRepo := _userRepo.NewUserMongoRepository(*mongoDatabase)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	_userHttp.NewUserHandler(router, userUsecase)

	appPort := fmt.Sprintf(":%v", bootstrap.App.Config.GetString("server.address"))
	log.Fatal(router.Run(appPort))
}
