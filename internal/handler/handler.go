package handler

import (
	"line-notification/common"
	"line-notification/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Ctx struct {
	*fiber.Ctx
	logger *zap.Logger
}

func (c *Ctx) Log() *zap.Logger {
	return c.logger
}

func (c *Ctx) Status(status int) *Ctx {
	ctx := c.Ctx.Status(status)
	return &Ctx{
		ctx,
		c.logger,
	}
}

func (c *Ctx) JSON(data interface{}) error {
	if c.Context().Response.StatusCode() > 299 {
		switch v := data.(type) {
		case map[string]string:
			c.logger.Error(v["error"])
		case *response.ErrResponse:
			c.logger.Error(v.Error)
		case error:
			c.logger.Error(v.Error())
		}
	}
	return c.Ctx.JSON(data)
}

type Handler func(*Ctx) error

func Helper(handler Handler, logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		xid := c.Request().Header.Peek(common.XRequestID)
		return handler(&Ctx{
			c,
			logger.With(zap.String(common.XRequestID, string(xid))),
		})
	}
}
