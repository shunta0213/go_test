package drivers

import "github.com/shunta0213/test_go_restapi/usecases/user"

type UserController struct {
	Interactor user.UserInteractor
}
