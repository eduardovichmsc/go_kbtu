package models

import "database/sql"

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Budget int    `json:"budget"`
}

type MovieModel struct {
	DB *sql.DB
}

func (m *MovieModel) GetAll() ([]Movie, error) {
	rows, err := m.DB.Query("SELECT id, title, genre, budget FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Budget); err != nil {
			continue
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (m *MovieModel) Insert(movie *Movie) error {
	query := "INSERT INTO movies(title, genre, budget) VALUES($1, $2, $3) RETURNING id"
	err := m.DB.QueryRow(query, movie.Title, movie.Genre, movie.Budget).Scan(&movie.ID)
	return err
}