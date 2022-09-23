package user

import (
	"context"

	"github.com/shunta0213/go_test/entities"
)

/*
Implement: Interactor, User: Controller
*/
type UserInputPort interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, id int) error
	GetRangeUsers(ctx context.Context, start_id int, end_id int) error
	GetUsers(ctx context.Context, id ...int) error
}

/*
Implement: Presenter, User: Interactor
*/
type UserOutputPort interface {
	AddUser(ctx context.Context, user *entities.User, err error) (*entities.User, error)
	OutPutUser(ctx context.Context, user *entities.User, err error) (*entities.User, error)
	OutputUsers(ctx context.Context, users *entities.Users, err error) (*entities.Users, error)
}

type UserRepository interface {
	AddUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUser(ctx context.Context, id int) (*entities.User, error)
	GetRangeUsers(ctx context.Context, start_id int, end_id int) (*entities.Users, error)
	GetUsers(ctx context.Context, id ...int) (*entities.Users, error)
}
