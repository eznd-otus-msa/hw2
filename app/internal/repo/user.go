package repo

import "github.com/eznd-otus-msa/hw2/app/internal/domain"

type UserReader interface {
    Get(int64) (*domain.User, error)
}

type UserCreator interface {
    Create(*domain.User) error
}

type UserUpdater interface {
    Update(int64, *domain.User) error
}

type UserPartialUpdater interface {
    PartialUpdate(int64, domain.UserPartialData) error
}

type UserDeleter interface {
    Delete(domain.UserId) error
}

type UserRepo interface {
    UserReader
    UserCreator
    UserUpdater
    UserPartialUpdater
    UserDeleter
}
