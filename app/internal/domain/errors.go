package domain

import "errors"

type HttpError interface {
	HttpCode() int
}

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidUserId = errors.New("invalid user id")
)
