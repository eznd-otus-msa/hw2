package http

import (
	"errors"
	"github.com/eznd-otus-msa/hw2/app/internal/domain"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type errBadRequest struct {
	error
}

func (e *errBadRequest) HttpCode() int {
	return fiber.StatusBadRequest
}
func (e *errBadRequest) Error() string {
	return "bad request"
}

var (
	ErrBadRequest       = errBadRequest{}
	ErrMethodNotAllowed = errors.New("method not allowed")
)

func json(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(data)
}

func fail(c *fiber.Ctx, err error) error {
	code := codeByErr(err)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": err.Error(),
	})
}

func codeByErr(err error) int {
	if errors.Is(err, ErrMethodNotAllowed) {
		return fiber.StatusMethodNotAllowed
	}

	if errors.Is(err, domain.ErrInvalidUserId) {
		return fiber.StatusBadRequest
	}
	if errors.Is(err, domain.ErrUserNotFound) {
		return fiber.StatusNotFound
	}

	if _, ok := err.(domain.HttpError); ok {
		return err.(domain.HttpError).HttpCode()
	}

	if _, ok := err.(validation.Errors); ok {
		return fiber.StatusBadRequest
	}
	return fiber.StatusInternalServerError
}

func DefaultResponse(ctx *fiber.Ctx) error {
	return fail(ctx, ErrMethodNotAllowed)
}
