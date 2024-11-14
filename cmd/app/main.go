package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-clean/config"
	"go-clean/internal/api/rest/middlewares"
	"go-clean/internal/composites"
	"go-clean/pkg/postgres"
	"go-clean/pkg/redis"
	"io"
	"os"
)

var ctx = context.Background()

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	cfg := initConfig()
	store := initRedis(cfg)
	defer func() {
		_ = store.Close()
	}()
	db := initPostgres(cfg)
	defer db.Close()
	initPostgresMigrations(db.ConnectionString)

	// Rest server
	restServer := echo.New()
	restServer.Logger.SetOutput(io.Discard)
	router := restServer.Group("/api")

	router.Use(middlewares.RequestIDMiddleware)
	router.Use(middlewares.StartLogging)
	router.Use(middlewares.ResponseMiddleware)

	composites.NewRestAuth(cfg, router, store, db)
	composites.NewRestUser(cfg, router, store, db)

	fmt.Println("Server started on port: " + cfg.Server.Port)
	if err := restServer.Start(":" + cfg.Server.Port); err != nil {
		panic(err)
	}
}

func initConfig() *config.Config {
	fmt.Println("Config initializing...")
	cfg, err := config.InitConfig("./env.yaml")
	if err != nil {
		panic("Config not initialized")
	}
	fmt.Println("Config initialized")
	return cfg
}

func initRedis(cfg *config.Config) *redis.Redis {
	fmt.Println("Redis connecting...")
	redisComposite, err := composites.NewRedis(ctx, &redis.Options{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		panic("Redis not connected")
	}
	fmt.Println("Redis connected")
	return redisComposite
}

func initPostgres(cfg *config.Config) *postgres.Postgres {
	fmt.Println("Postgres connecting...")
	postgresComposite, err := composites.NewPostgres(ctx, &postgres.Options{
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
	fmt.Println("Postgres connected")
	return postgresComposite
}
