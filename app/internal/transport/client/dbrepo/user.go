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

func (r *userRepo) Get(id domain.UserId) (u *domain.User, err error) {
	err = r.db.Model(u).Where(id).First(&u).Error
	return
}

func (r *userRepo) Exists(id domain.UserId) (exists bool, err error) {
	err = r.db.Model(&domain.User{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error
	return
}

func (r *userRepo) Create(u *domain.User) (*domain.User, error) {
	err := r.db.Create(u).Error
	return u, err
}

func (r *userRepo) Update(id domain.UserId, user *domain.User) (*domain.User, error) {
	user.Id = id
	err := r.db.Save(user).Error
	return user, err
}

func (r *userRepo) PartialUpdate(id domain.UserId, data *domain.UserPartialData) (u *domain.User, err error) {
	upd := make(map[string]interface{})
	for k, v := range data.All() {
		upd[k] = v
	}

	err = r.db.Model(&domain.User{}).Where(id).Updates(upd).Error
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *userRepo) Delete(id domain.UserId) error {
	return r.db.Delete(&domain.User{}, id).Error
}
