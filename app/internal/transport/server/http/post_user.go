package http

//
//import (
//    "github.com/eznd-otus-msa/hw2/app/internal/domain"
//    "github.com/gofiber/fiber/v2"
//)
//
//func NewPostUser(c domain.UserCreator) *postUserHandler {
//    return &postUserHandler{
//        creator: c,
//    }
//}
//
//type postUserHandler struct {
//    creator domain.UserCreator
//}
//
//func (h *postUserHandler) Handle() fiber.Handler {
//    return func(ctx *fiber.Ctx) error {
//        var u domain.User
//        err := ctx.BodyParser(&u)
//        if err != nil {
//            return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad payload"})
//        }
//
//        err = h.creator.Create(&u)
//        if err != nil {
//            return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user: " + err.Error()})
//        }
//        return ctx.Status(fiber.StatusOK).JSON(u)
//    }
//}
