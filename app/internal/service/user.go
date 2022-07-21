package service

import (
	"github.com/eznd-otus-msa/hw2/app/internal/domain"
	"github.com/eznd-otus-msa/hw2/app/internal/repo"
	"github.com/eznd-otus-msa/hw2/app/pkg/nullable"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserData struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserCreate UserData

func (u UserCreate) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(2, 32)),
		validation.Field(&u.FirstName, validation.Required, validation.Length(1, 32)),
		validation.Field(&u.LastName, validation.Required, validation.Length(1, 32)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Phone, validation.Required, is.E164),
	)
}

func (u UserCreate) ToDomain() *domain.User {
	return &domain.User{
		Id:        0,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}

type UserUpdate UserData

type UserPartialUpdate struct {
	Username  nullable.String `json:"username"`
	FirstName nullable.String `json:"firstName"`
	LastName  nullable.String `json:"lastName"`
	Email     nullable.String `json:"email"`
	Phone     nullable.String `json:"phone"`
}

type UserReader interface {
	Get(domain.UserId) (*domain.User, error)
}

type UserCreator interface {
	Create(*UserCreate) error
}

type UserUpdater interface {
	Update(domain.UserId, *domain.User) error
}

type UserPartialUpdater interface {
	PartialUpdate(domain.UserId, domain.UserPartialData) error
}

type UserDeleter interface {
	Delete(domain.UserId) error
}

type UserService interface {
	UserReader
	UserCreator
	UserUpdater
	UserPartialUpdater
	UserDeleter
}

type userService struct {
	reader         repo.UserReader
	creator        repo.UserCreator
	updater        repo.UserUpdater
	partialUpdater repo.UserPartialUpdater
	deleter        repo.UserDeleter
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{
		reader:         repo,
		creator:        repo,
		updater:        repo,
		partialUpdater: repo,
		deleter:        repo,
	}
}

func (s *userService) Get(id domain.UserId) (*domain.User, error) {
	err := id.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.reader.Get(int64(id))
	if err != nil {
		if err.Error() == "record not found" {
			return nil, domain.ErrUserNotFound
		}
	}
	return user, err
}

func (s *userService) Create(user *UserCreate) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	return s.creator.Create(user.ToDomain())
}

func (s *userService) Update(id domain.UserId, user *domain.User) error {
	return s.updater.Update(int64(id), user)
}

func (s *userService) PartialUpdate(id domain.UserId, data domain.UserPartialData) error {
	return s.partialUpdater.PartialUpdate(int64(id), data)
}

func (s *userService) Delete(id domain.UserId) error {
	return s.deleter.Delete(id)
}
