package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"practice_7/internal/handler"
	"practice_7/internal/repository"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:admin@localhost:5432/practice_7?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	repo := repository.NewRepository(db)
	h := handler.NewHandler(repo)

	http.HandleFunc("/users", h.GetUsersHandler)
	http.HandleFunc("/users/common-friends", h.GetCommonFriendsHandler)

	fmt.Println("Server is running on port localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}