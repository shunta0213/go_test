package controllers

import (
	"context"

	"github.com/shunta0213/go_test/usecases/user"
)

type UserController struct {
	Interactor user.UserInteractor
}

func (controller *UserController) AddUser(ctx context.Context)
