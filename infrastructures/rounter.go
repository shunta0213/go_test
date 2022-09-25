package infrastructures

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shunta0213/go_test/adapters/controllers"
	"github.com/shunta0213/go_test/infrastructures/database"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	/*
		Users
	*/
	userController := controllers.NewUserController(*database.NewSqlHandler())
	router.GET("/users/all", func(ctx *gin.Context) { userController.GetAllUser(context.TODO(), ctx) })
	router.GET("/users/range", func(ctx *gin.Context) { userController.GetRangeUsers(context.TODO(), ctx) })
	router.GET("/users", func(ctx *gin.Context) { userController.GetUsers(context.TODO(), ctx) })
	router.GET("/user/:id", func(ctx *gin.Context) { userController.GetUser(context.TODO(), ctx) })
	router.POST("/user", func(ctx *gin.Context) { userController.AddUser(context.TODO(), ctx) })

	Router = router
}
