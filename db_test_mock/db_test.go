package dbtestmock

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DbSuite struct {
	suite.Suite
	DB   *gorm.DB
	Mock sqlmock.Sqlmock

	Repository UserRepository
	User       User
}

func (s *DbSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, s.Mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.Repository = UserRepository{DB: s.DB}
}

func (s *DbSuite) Test_repository_Get() {
	var (
		id, _ = uuid.NewV4()
		name  = "test-name"
	)

	s.Mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (id = ?)")).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(id.String(), name))

	res, err := s.Repository.Get(id)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), &User{ID: id, Name: name}, res)
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(DbSuite))
}
