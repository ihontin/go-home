package main

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"studentgit.kata.academy/Alkolex/go-kata/course3/1.reflection/1.reflection_using/task3.1.1.4/dao"
	"testing"
)

func TestDAO_BuildSelect(t *testing.T) {
	// Создание мока для базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}
	dbx := sqlx.NewDb(db, "sqlite3")
	d := dao.NewDAO(dbx)
	condition := dao.Condition{
		LimitOffset: &dao.LimitOffset{
			Offset: 5,
			Limit:  3,
		},
		Equal: map[string]interface{}{
			"id": 1,
		},
		Order: []*dao.Order{
			{
				Field: "id",
				Asc:   true,
			},
		},
	}
	expect := "SELECT id, username FROM"
	expectedQuery := expect + " users WHERE id = $1 ORDER BY id ASC LIMIT 3 OFFSET 5"
	query, _, err := d.BuildSelect("users", condition, "id", "username")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Проверка, что возвращенный SQL-запрос соответствует ожидаемому
	if query != expectedQuery {
		t.Errorf("Expected query %v, but got %v", expectedQuery, query)
	}
	// Проверка, что все ожидаемые SQL-запросы были выполнены
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %v", err)
	}
	defer db.Close()
	dbx := sqlx.NewDb(db, "sqlite3")
	d := dao.NewDAO(dbx)

	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "john.doe@example.com",
		Address:   "123 Main St",
		Status:    1,
		DeletedAt: "",
	}
	intoUsers := "INTO users"
	mock.ExpectExec("INSERT "+intoUsers).
		WithArgs(user.ID, user.FirstName, user.LastName, user.Username, user.Email, user.Address, user.Status, user.DeletedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = d.Create(context.Background(), &user)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "mock")
	d := dao.NewDAO(dbx)
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "john.doe@example.com",
		Address:   "123 Main St",
		Status:    1,
		DeletedAt: "",
	}

	selected := "SELECT"
	mock.ExpectQuery(selected + " (.+) FROM (.+) WHERE (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "address", "status", "deleted_at"}).
			AddRow(user.ID, user.FirstName, user.LastName, user.Username, user.Email, user.Address, user.Status, user.DeletedAt))

	var dest []User
	err = d.List(context.Background(), &dest, &User{}, dao.Condition{
		Equal: map[string]interface{}{
			"id": 1,
		},
	})

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

	expected := []User{user}
	assert.Equal(t, expected, dest)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "mock")
	d := dao.NewDAO(dbx)
	updated := "UPDATE"
	mock.ExpectQuery("SELECT EXISTS(.+)").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
	mock.ExpectExec(updated + " (.+) SET (.+) WHERE (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

	entity := &User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Address:   "123 Main St",
		Status:    1,
		DeletedAt: "",
	}

	condition := dao.Condition{
		Equal: map[string]interface{}{
			"id": 1,
		},
	}

	err = d.Update(context.Background(), entity, condition)

	assert.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
