package test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shunta0213/go_test/adapters/gateways"
	"github.com/shunta0213/go_test/entities"
	"github.com/shunta0213/go_test/infrastructures/database"
	"github.com/stretchr/testify/assert"
)

func newDummySqlHandler(db *sql.DB) database.SqlHandler {
	sqlHandler := new(database.SqlHandler)
	sqlHandler.Conn = db
	return *sqlHandler
}

func TestGetUser(t *testing.T) {

	// Table for test

	table := []struct {
		testName string
		id       int
		user     entities.User
		err      error
	}{
		{
			"GetUserById1",
			1,
			entities.User{
				Id:        1,
				Username:  "User",
				Firstname: "first",
				Lastname:  "last",
				Email:     "email",
				Contact:   "contact",
			},
			nil,
		},
	}

	// db is for connecting mock db.
	// mock is for setting mock.
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Error("SqlMock not Work")
	}

	defer db.Close()

	repo := &gateways.UserGateway{
		SqlHandler: newDummySqlHandler(db),
	}

	query := "SELECT * FROM users WHERE id=?"

	for _, tt := range table {
		t.Run(tt.testName, func(t *testing.T) {

			b := tt.user

			rows := sqlmock.NewRows([]string{
				"id", "username", "firstname", "lastname", "email", "contact",
			}).AddRow(b.Id, b.Username, b.Firstname, b.Lastname, b.Email, b.Contact)

			// ExpectQuery mocks the combination of expected SELECT statement and retrieved result.
			mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1, tt.id).WillReturnRows(rows)

			got, err := repo.GetUser(context.TODO(), b.Id)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, b, got)
		})
	}

}
