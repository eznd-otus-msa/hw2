package http

import (
    "errors"
    "github.com/eznd-otus-msa/hw2/app/internal/domain"
    "github.com/gofiber/fiber/v2"
    "net/http"
)

var (
    ErrBadRequest       = errors.New("bad request")
    ErrMethodNotAllowed = errors.New("method not allowed")
)

func httpResponse(c *fiber.Ctx, data interface{}) error {
    return c.Status(http.StatusOK).JSON(data)
}

func httpError(c *fiber.Ctx, err error) error {
    code := codeByErr(err)
    return c.Status(code).JSON(fiber.Map{
        "code":    code,
        "message": err.Error(),
    })
}

func codeByErr(err error) int {
    if errors.Is(err, ErrBadRequest) {
        return fiber.StatusBadRequest
    }
    if errors.Is(err, ErrMethodNotAllowed) {
        return fiber.StatusMethodNotAllowed
    }

    if errors.Is(err, domain.ErrInvalidUserId) {
        return fiber.StatusBadRequest
    }
    if errors.Is(err, domain.ErrUserNotFound) {
        return fiber.StatusNotFound
    }
    return fiber.StatusInternalServerError
}

func DefaultResponse(ctx *fiber.Ctx) error {
    return httpError(ctx, ErrMethodNotAllowed)
}
