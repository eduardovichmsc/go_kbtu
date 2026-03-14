package repository

import (
	"database/sql"
	"fmt"
	"practice_7/internal/models"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetPaginatedUsers(page, pageSize int, filters map[string]string, orderBy string) (models.PaginatedResponse, error) {
	offset := (page - 1) * pageSize
	args := []interface{}{}
	conditions := []string{}
	argId := 1

	for key, val := range filters {
		if val != "" {
			conditions = append(conditions, fmt.Sprintf("%s = $%d", key, argId))
			args = append(args, val)
			argId++
		}
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	var totalCount int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM users %s`, whereClause)
	err := r.db.QueryRow(countQuery, args...).Scan(&totalCount)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	allowedOrderFields := map[string]bool{"id": true, "name": true, "email": true, "gender": true, "birth_date": true}
	if !allowedOrderFields[orderBy] {
		orderBy = "id" // Default case
	}

	query := fmt.Sprintf(`SELECT id, name, email, gender, birth_date FROM users %s ORDER BY %s LIMIT $%d OFFSET $%d`, whereClause, orderBy, argId, argId+1)
	args = append(args, pageSize, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return models.PaginatedResponse{}, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return models.PaginatedResponse{}, err
		}
		users = append(users, u)
	}

	if users == nil {
		users = []models.User{}
	}

	return models.PaginatedResponse{
		Data:       users,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

func (r *Repository) GetCommonFriends(userID1, userID2 string) ([]models.User, error) {
	query := `
		SELECT u.id, u.name, u.email, u.gender, u.birth_date
		FROM users u
		JOIN user_friends uf1 ON u.id = uf1.friend_id
		JOIN user_friends uf2 ON u.id = uf2.friend_id
		WHERE uf1.user_id = $1 AND uf2.user_id = $2
	`

	rows, err := r.db.Query(query, userID1, userID2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	friends := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return nil, err
		}
		friends = append(friends, u)
	}

	if friends == nil {
		friends = []models.User{}
	}

	return friends, nil
}