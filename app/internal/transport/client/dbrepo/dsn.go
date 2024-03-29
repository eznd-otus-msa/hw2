package dbrepo

import (
    "fmt"
    "github.com/eznd-otus-msa/hw2/app/internal/domain"
)

func Dsn(cfg domain.Db) string {
    connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
        cfg.Host,
        cfg.Port,
        cfg.User,
        cfg.DbName,
        cfg.Password,
    )
    if cfg.Extras != "" {
        connString += " " + cfg.Extras
    }
    return connString
}
