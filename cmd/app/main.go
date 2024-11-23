package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-clean/config"
	"go-clean/internal/api/rest/middlewares"
	"go-clean/internal/composites"
	"go-clean/migrations"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
	"io"
	"os"
)

var ctx = context.Background()

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	// Config
	cfg, err := config.InitConfig("./env.yaml")
	if err != nil {
		panic("Config not initialized")
	}

	// Redis
	store, err := composites.NewRedis(ctx, &redis.Options{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		panic("Redis not connected")
	}
	defer func() {
		_ = store.Close()
	}()

	// Postgres
	db, err := composites.NewPostgres(context.Background(), &postgres.Options{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		panic("Postgres not connected")
	}
	defer db.Close()

	// Migrations
	migrations.Run(&migrations.Options{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	// Rest server
	restServer := echo.New()
	restServer.Logger.SetOutput(io.Discard)
	router := restServer.Group("/api")

	router.Use(middlewares.StartLogging)
	router.Use(middlewares.RequestIDMiddleware)
	router.Use(middlewares.ResponseMiddleware)

	composites.NewRestAuth(cfg, router, store, db)
	composites.NewRestUser(cfg, router, store, db)
	composites.NewGroups(cfg, router, store, db)
	composites.NewGroupUsers(cfg, router, store, db)
	composites.NewGroupNote(cfg, router, store, db)

	fmt.Println("Server started on port: " + cfg.Server.Port)
	if err = restServer.Start(":" + cfg.Server.Port); err != nil {
		panic(err)
	}
}
