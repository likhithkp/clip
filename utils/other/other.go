package other

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseStruct struct{}

func NewResponseStruct() *ResponseStruct {
	return &ResponseStruct{}
}

type Response struct {
	Success bool
	Message string
	Data    any
}

func (responseStruct *ResponseStruct) Response(ctx *fiber.Ctx, status int, success bool, message string, data any) error {
	return ctx.Status(status).JSON(Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}
