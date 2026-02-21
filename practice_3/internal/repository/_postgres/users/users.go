package users

import (
	"database/sql"
	"errors"
	"fmt"
	"practice_3/internal/pkg/modules"
	"practice_3/internal/repository/_postgres"
	"time"
)

type Repository struct {
	db               *_postgres.Dialect
	executionTimeout time.Duration
}

func NewUserRepository(db *_postgres.Dialect) *Repository {
	return &Repository{
		db:               db,
		executionTimeout: time.Second * 5,
	}
}

func (r *Repository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*modules.User, error) {
	var user modules.User
	err := r.db.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *modules.User) (int, error) {
	var newID int
	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.DB.QueryRow(query, user.Name, user.Email, user.Age).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func (r *Repository) UpdateUser(id int, user *modules.User) error {
	query := `UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4`
	res, err := r.db.DB.Exec(query, user.Name, user.Email, user.Age, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cannot update: user with ID %d does not exist", id)
	}

	return nil
}

func (r *Repository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	res, err := r.db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("cannot delete: user with ID %d does not exist", id)
	}

	return nil
}