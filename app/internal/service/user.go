package service

import (
    "github.com/eznd-otus-msa/hw2/app/internal/domain"
    "github.com/eznd-otus-msa/hw2/app/internal/repo"
)

type UserReader interface {
    Get(domain.UserId) (*domain.User, error)
}

type UserCreator interface {
    Create(*domain.User) error
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

func (s *userService) Create(user *domain.User) error {
    return s.creator.Create(user)
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
