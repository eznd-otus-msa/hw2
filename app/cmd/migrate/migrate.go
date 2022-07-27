package main

import (
	"fmt"
	"github.com/ezn-go/mixture"
	"github.com/eznd-otus-msa/hw2/app/internal/transport/client/dbrepo"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jinzhu/configor"

	"github.com/eznd-otus-msa/hw2/app/internal/domain"
	_ "github.com/eznd-otus-msa/hw2/app/migrations"
)

func main() {
	var cfg domain.Config
	err := configor.Load(&cfg, "config/config.yaml", "./config.yaml")
	if err != nil {
		panic(err)
	}

	/*
	   DB_HOST=localhost
	   DB_NAME=postgres
	   DB_PASSWORD=example
	   DB_PORT=5432
	   DB_USER=postgres
	   DB_EXTRAS=sslmode=disable TimeZone=Europe/Moscow client_encoding=UTF8
	*/
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbrepo.Dsn(cfg.Db),
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = mixture.Apply(db, cfg.App.Env)
	if err != nil {
		panic(err)
	}

	fmt.Println("migrations applied")
}
