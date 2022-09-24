package controllers

import (
	"context"
	"strconv"

	"github.com/shunta0213/go_test/adapters/gateways"
	"github.com/shunta0213/go_test/adapters/presenters"
	"github.com/shunta0213/go_test/entities"
	"github.com/shunta0213/go_test/infrastructures/database"
	"github.com/shunta0213/go_test/usecases/user"
)

type UserController struct {
	Interactor user.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: user.UserInteractor{
			OutputPort: &presenters.UserPresenter{},
			Repository: &gateways.UserGateway{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) AddUser(ctx context.Context, c Context) {
	u := entities.User{}
	c.Bind(&u)
	user, err := controller.Interactor.AddUser(ctx, &u)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(201, user)
}

func (controller *UserController) GetUser(ctx context.Context, c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.Repository.GetUser(ctx, id)

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, user)
}
