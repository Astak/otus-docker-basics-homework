package data

import (
	"context"
	"errors"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(id int64) (*models.User, error)
	DeleteUser(id int64) error
	UpdateUser(id int64, user *models.User) (*models.User, error)
}

type PostgresUserRepository struct {
	db *Database
}

func NewUserRepository(db *Database) (*PostgresUserRepository, error) {
	return &PostgresUserRepository{db: db}, nil
}

func (repo PostgresUserRepository) CreateUser(user *models.User) (*models.User, error) {
	sql := `
	INSERT INTO users (username, firstname, lastname, email, phone)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	var id int64
	err := repo.db.Conn.QueryRow(context.Background(), sql, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone).Scan(&id)
	if err != nil {
		panic(err)
	}
	user.ID = id
	return user, nil
}

func (repo PostgresUserRepository) GetUser(id int64) (*models.User, error) {
	sql := `
	SELECT id, username, firstname, lastname, email, phone
	FROM users
	WHERE id = $1
	`
	var user models.User
	rows, err := repo.db.Conn.Query(context.Background(), sql, id)
	if err != nil {
		panic(err)
	}
	err = pgxscan.ScanOne(&user, rows)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		panic(err)
	}
	return &user, nil
}

func (repo PostgresUserRepository) DeleteUser(id int64) error {
	sql := `
	DELETE FROM users
	WHERE id = $1
	`
	_, err := repo.db.Conn.Exec(context.Background(), sql, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (repo PostgresUserRepository) UpdateUser(id int64, user *models.User) (*models.User, error) {
	sql := `
	UPDATE users SET
		username = $2,
		firstname = $3,
		lastname = $4,
		email = $5,
		phone = $6
	WHERE id = $1
	`
	res, err := repo.db.Conn.Exec(context.Background(), sql, id, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil {
		panic(err)
	}
	count := res.RowsAffected()
	if count == 0 {
		return nil, nil
	}
	user.ID = id
	return user, nil
}
