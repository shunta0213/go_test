package presenters

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/shunta0213/go_test/entities"
	"github.com/shunta0213/go_test/models"
)

type UserPresenter struct{}

func (u *UserPresenter) AddUser(ctx context.Context, user *models.User, err error) (*entities.User, error) {
	entities_user := &entities.User{}
	copier.Copy(entities_user, user)
	return entities_user, err
}

func OutPutUser(ctx context.Context, user *models.User, err error) (*entities.User, error) {
	entities_user := &entities.User{}
	copier.Copy(entities_user, user)
	return entities_user, err
}

// Im not sure this works.
func (u *UserPresenter) OutputUsers(ctx context.Context, users []*models.User, err error) (*entities.Users, error) {
	entities_users := &entities.Users{}
	copier.Copy(entities_users, users)
	return entities_users, err
}
