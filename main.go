package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// For sql, postgres driver

	_ "github.com/lib/pq"
	"github.com/shunta0213/go_test/database"
	"github.com/shunta0213/go_test/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {

	// Create User
	// d := `{"username": "test4", "firstname": "first4", "lastname": "last4", "email":"test4@test4.io"}`
	// got, err := CreateUser(context.TODO(), []byte(d), database.Db)
	// if got != nil {
	// 	fmt.Printf("Done")
	// }

	// necessary
	router := gin.Default()
	users, err := ListUsers(context.TODO(), database.Db)
	router.GET("/user", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, users)
	})

	if err != nil {
		fmt.Println(err)
	}

	router.Run("localhost:8080")
}

func CreateUser(ctx context.Context, b []byte, db *sql.DB) (*models.User, error) {
	u := &models.User{}
	if err := json.Unmarshal(b, u); err != nil {
		return nil, err
	}
	err := u.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return u, err
}

func ListUsers(ctx context.Context, db *sql.DB) ([]*models.User, error) {
	users, err := models.Users(
		qm.Where("id=?", 1),
	).All(ctx, db)
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
