package app

import (
	"context"
	"log"
	"net/http"
	"practice_3/internal/handler"
	"practice_3/internal/middleware"
	"practice_3/internal/pkg/modules"
	"practice_3/internal/repository"
	"practice_3/internal/repository/_postgres"
	"practice_3/internal/usecase"
	"time"
)

func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        "localhost",
		Port:        "5432",
		Username:    "postgres",
		Password:    "admin", 
		DBName:      "go_practice3",     
		SSLMode:     "disable",
		ExecTimeout: 5 * time.Second,
	}
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := initPostgreConfig()
	pgxDialect := _postgres.NewPGXDialect(ctx, dbConfig)

	repos := repository.NewRepositories(pgxDialect)
	userUsecase := usecase.NewUserUsecase(repos.UserRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", userHandler.Healthcheck)
	mux.HandleFunc("GET /users", userHandler.GetAll)
	mux.HandleFunc("GET /users/{id}", userHandler.GetByID)
	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("PUT /users/{id}", userHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", userHandler.Delete)

	handlerWithAuth := middleware.AuthMiddleware(mux)
	finalHandler := middleware.LoggingMiddleware(handlerWithAuth)

	log.Println("Server is running on :8080...")
	if err := http.ListenAndServe("localhost:8080", finalHandler); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}