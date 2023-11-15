package mysql

import (
	"context"
	"errors"
	mailing "mail_service/internal"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CourseRepository_Save_RepositoryError(t *testing.T) {
	courseID, courseName, courseDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Course", "37a0f027-15e6-47cc-a5d2-64183281087e"
	course, err := mailing.NewCliente(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO clientes (id, nombre, id_plan) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnError(errors.New("something-failed"))

	repo := NewClienteRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CourseRepository_Save_Succeed(t *testing.T) {
	courseID, courseName, courseDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Course", "37a0f027-15e6-47cc-a5d2-64183281087e"
	course, err := mailing.NewCliente(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO clientes (id, nombre, id_plan) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewClienteRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
