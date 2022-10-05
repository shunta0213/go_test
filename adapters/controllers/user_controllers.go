package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/shunta0213/test_go_restapi/adapters/gateways"
	"github.com/shunta0213/test_go_restapi/adapters/presenters"
	"github.com/shunta0213/test_go_restapi/entities"
	"github.com/shunta0213/test_go_restapi/infrastructures/database"
	"github.com/shunta0213/test_go_restapi/usecases/user"
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

func (controller *UserController) AddUser(ctx context.Context, c GinContext) {
	u := entities.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Repository.AddUser(ctx, &u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) GetUser(ctx context.Context, c GinContext) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.Repository.GetUser(ctx, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) GetRangeUsers(ctx context.Context, c GinContext) {
	startId, _ := strconv.Atoi(c.Query("start"))
	endId, _ := strconv.Atoi(c.Query("end"))
	users, err := controller.Interactor.Repository.GetRangeUsers(ctx, startId, endId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, users)
}

func (controller *UserController) GetAllUser(ctx context.Context, c GinContext) {
	users, err := controller.Interactor.Repository.GetAllUser(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, users)
}

type GetUsersPathRequest struct {
	Ids []int `form:"ids[]"`
}

func (controller *UserController) GetUsers(ctx context.Context, c GinContext) {
	var req GetUsersPathRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	ids := req.Ids
	users, err := controller.Interactor.Repository.GetUsers(ctx, ids...)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, users)
}
