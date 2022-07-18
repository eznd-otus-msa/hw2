package dbrepo

import (
    "github.com/eznd-otus-msa/hw2/app/internal/domain"
    "github.com/eznd-otus-msa/hw2/app/internal/repo"
    "gorm.io/gorm"
)

type userRepo struct {
    db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
    return &userRepo{db: db}
}

func (r *userRepo) Get(id int64) (*domain.User, error) {
    var u domain.User
    err := r.db.Model(u).Where(id).First(&u).Error
    return &u, err
}

func (r *userRepo) Create(u *domain.User) error {
    err := r.db.Create(u).Error
    return err
}

func (r *userRepo) Update(id int64, user *domain.User) error {
    //TODO implement me
    panic("implement me")
}

func (r *userRepo) PartialUpdate(id int64, data domain.UserPartialData) error {
    //TODO implement me
    panic("implement me")
}

func (r *userRepo) Delete(id domain.UserId) error {
    //TODO implement me
    panic("implement me")
}
