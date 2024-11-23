package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/VadimSmD/KR/gen/oas"
	"github.com/VadimSmD/KR/internal/controller"
	"github.com/VadimSmD/KR/internal/repo"
	"github.com/VadimSmD/KR/internal/service"
)

type app struct {
	server *http.Server
}

func newApp(ctx context.Context, env environment) (*app, error) {
	pg, err := pgxpool.New(ctx, env.Postgres.DSN)
	if err != nil {
		return nil, fmt.Errorf("newApp: %w", err)
	}
	userRepo := repo.NewUserRepo(pg)
	userService := service.NewUserService(userRepo)
	userHandler := controller.NewUserController(userService)

	handler := &Handler{
		UserHandler: userHandler.UserHandler,
	}
	ogenServer, err := oas.NewServer(handler)
	if err != nil {
		return nil, fmt.Errorf("newApp: %w", err)
	}

	httpServer := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
		Handler:           ogenServer,
	}
	return &app{
		server: httpServer,
	}, nil
}

func (a *app) run(ctx context.Context) <-chan struct{} {
	go func() {
		log.Println("Starting HTTP server at :8080")
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("app run: %s", err)
			return
		}
		log.Println("HTTP server gracefully stopped")
	}()
	return ctx.Done()
}

func (a *app) shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server")
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
