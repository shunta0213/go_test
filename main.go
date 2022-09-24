package main

import (
	"context"
	"database/sql"

	// For sql, postgres driver

	"github.com/shunta0213/go_test/infrastructures"
	"github.com/shunta0213/go_test/models"
)

func main() {

	// Create User
	// user := entities.User{Username: "user5", Firstname: "first5", Lastname: "last5", Email: "first5@last5.io", Contact: "050"}
	// gateways.AddUser(context.TODO(), &user)

	// necessary
	// router := gin.Default()
	// users, err := gateways.GetRangeUsers(context.TODO(), 1, 3)
	// router.GET("/user", func(ctx *gin.Context) {
	// 	ctx.IndentedJSON(http.StatusOK, users)
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// router.Run("localhost:8080")

	infrastructures.Router.Run()

}

func ListUsers(ctx context.Context, db *sql.DB) ([]*models.User, error) {
	users, err := models.Users().All(ctx, db)
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []*models.User{}
	}

	return users, nil
}

func ListUser(ctx context.Context, db *sql.DB) (*models.User, error) {
	user, err := models.Users().One(ctx, db)

	if err != nil {
		return nil, err
	}
	return user, nil
}
