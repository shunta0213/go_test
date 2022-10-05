package gateways

import (
	"context"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/shunta0213/test_go_restapi/entities"
	"github.com/shunta0213/test_go_restapi/infrastructures/database"
	"github.com/shunta0213/test_go_restapi/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

/*
This contains Database information which is in the infra layer,
so this breaks dependency inversion principal.
To solve this problem, Add sql handler to this gateways (adaptors) layer.
*/

type UserGateway struct {
	database.SqlHandler
}

/*
When create user
*/
func (gateway *UserGateway) AddUser(ctx context.Context, user *entities.User) (*models.User, error) {

	u := models.User{}
	err := copier.Copy(&u, user)

	if err != nil {
		return nil, err
	}

	err = u.Insert(ctx, gateway.Conn, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &u, err
}

/*
When you get a user by ID
*/
func (gateway *UserGateway) GetUser(ctx context.Context, id int) (*models.User, error) {

	user, err := models.FindUser(ctx, gateway.Conn, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		user = &models.User{}
	}

	return user, nil

}

/*
When you get users by Ids.
*/
func (gateway *UserGateway) GetUsers(ctx context.Context, ids ...int) ([]*models.User, error) {

	// make Query
	// maybe can improve speed
	var query string
	for r, i := range ids {
		if r == len(ids)-1 {
			query += ("id=" + strconv.Itoa(i))
			continue
		}
		query += ("id=" + strconv.Itoa(i) + " OR ")
	}

	users, err := models.Users(
		qm.Where(query),
	).All(ctx, gateway.Conn)

	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []*models.User{}
	}

	return users, nil
}

/*
When you get users in a certain range
*/

func (gateways *UserGateway) GetRangeUsers(ctx context.Context, start_id int, end_id int) ([]*models.User, error) {

	// make Query
	var query string
	for i := start_id; i <= end_id; i++ {
		if i == end_id {
			query += ("id=" + strconv.Itoa(i))
			continue
		}
		query += ("id=" + strconv.Itoa(i) + " OR ")
	}

	users, err := models.Users(
		qm.Where(query),
	).All(ctx, gateways.Conn)

	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []*models.User{}
	}

	return users, nil
}

/*
When you get all users
*/

func (gateways *UserGateway) GetAllUser(ctx context.Context) ([]*models.User, error) {

	users, err := models.Users().All(ctx, gateways.Conn)

	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []*models.User{}
	}

	return users, nil
}
