package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dblogger "gorm.io/gorm/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/configor"

	"github.com/eznd-otus-msa/hw2/app/internal/domain"
	"github.com/eznd-otus-msa/hw2/app/internal/service"
	"github.com/eznd-otus-msa/hw2/app/internal/transport/client/dbrepo"
	"github.com/eznd-otus-msa/hw2/app/internal/transport/server/http"
	_ "github.com/eznd-otus-msa/hw2/app/migrations"
)

func main() {
	var cfg domain.Config
	err := configor.Load(&cfg)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbrepo.Dsn(cfg.Db),
	}), &gorm.Config{
		Logger: dblogger.Default.LogMode(dblogger.Silent),
	})
	if err != nil {
		panic(err)
	}

	userRepo := dbrepo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	getUserHandler := http.NewGetUser(userService)
	postUserHandler := http.NewPostUser(userService)

	srv := fiber.New(fiber.Config{})
	srv.Use(logger.New())

	api := srv.Group("/api")

	api.Get("/user/:id", getUserHandler.Handle())
	api.Get("/users/:id", getUserHandler.Handle())

	api.Post("/user", postUserHandler.Handle())
	api.Post("/users", postUserHandler.Handle())

	srv.All("/*", http.DefaultResponse)

	err = srv.Listen(fmt.Sprintf(":%s", cfg.Http.Port))
	if err != nil {
		panic(err)
	}
}
