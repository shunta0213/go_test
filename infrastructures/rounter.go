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

	userController := controllers.NewUserController(*database.NewSqlHandler())

	router.GET("/user/:id", func(c *gin.Context) { userController.GetUser(context.TODO(), c) })

	Router = router
}
