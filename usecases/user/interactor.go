package user

import (
	"context"

	"github.com/shunta0213/go_test/entities"
)

type UserInteractor struct {
	OutputPort UserOutputPort
	Repository UserRepository
}

func (u *UserInteractor) AddUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	user_repo, err := u.Repository.AddUser(ctx, user)
	if err != nil {
		return u.OutputPort.AddUser(ctx, nil, err)
	}
	return u.OutputPort.AddUser(ctx, user_repo, nil)
}

func (u *UserInteractor) GetUser(ctx context.Context, id int) (*entities.User, error) {
	user, err := u.Repository.GetUser(ctx, id)
	if err != nil {
		return u.OutputPort.OutPutUser(ctx, nil, err)
	}

	return u.OutputPort.OutPutUser(ctx, user, nil)
}

func (u *UserInteractor) GetRangeUsers(ctx context.Context, start_id int, end_id int) (*entities.Users, error) {
	users, err := u.Repository.GetRangeUsers(ctx, start_id, end_id)
	if err != nil {
		return u.OutputPort.OutputUsers(ctx, users, err)
	}

	return u.OutputPort.OutputUsers(ctx, users, nil)
}

func (u *UserInteractor) GetUsers(ctx context.Context, ids ...int) (*entities.Users, error) {
	users, err := u.Repository.GetUsers(ctx, ids...)
	if err != nil {
		return u.OutputPort.OutputUsers(ctx, nil, err)
	}
	return u.OutputPort.OutputUsers(ctx, users, nil)
}
