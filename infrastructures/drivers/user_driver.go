package drivers

import "github.com/shunta0213/go_test/usecases/user"

type UserController struct {
	Interactor user.UserInteractor
}
