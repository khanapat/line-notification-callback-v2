package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"line-notification/common"
	"line-notification/response"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type middleware struct {
	ZapLogger *zap.Logger
}

func NewMiddleware(zapLogger *zap.Logger) *middleware {
	return &middleware{
		ZapLogger: zapLogger,
	}
}

func (m *middleware) JSONMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts(common.ApplicationJSON)
		return c.Next()
	}
}

func (m *middleware) ContextLocaleMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(common.LocaleKey, c.Query(common.LocaleKey))
		return c.Next()
	}
}

func (m *middleware) LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		if c.Request().Header.Peek(common.XRequestID) == nil {
			c.Request().Header.Add(common.XRequestID, uuid.New().String())
		}

		logger := m.ZapLogger.With(zap.String(common.XRequestID, string(c.Request().Header.Peek(common.XRequestID))))

		logger.Debug(common.RequestInfoMsg,
			zap.String("method", string(c.Request().Header.Method())),
			zap.String("host", string(c.Request().Header.Host())),
			zap.String("path_uri", c.Request().URI().String()),
			zap.String("remote_addr", c.Context().RemoteAddr().String()),
			zap.String("body", string(c.Request().Body())),
		)

		if err := c.Next(); err != nil {
			return err
		}
		logger.Debug(common.ResponseInfoMsg,
			zap.String("body", string(c.Response().Body())),
		)
		logger.Info("Summary Information",
			zap.String("method", string(c.Request().Header.Method())),
			zap.String("path_uri", c.Request().URI().String()),
			zap.Duration("duration", time.Since(start)),
			zap.Int("status_code", c.Response().StatusCode()),
		)
		return nil
	}
}

func (m *middleware) LineAuthenticationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Request().Header.Peek(common.LineSignatureHeader) == nil {
			err := errors.Wrapf(errors.New(fmt.Sprintf("Token doesn't exist.")), response.ValidateAuthenticationTokenError)
			return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).AuthenticationTokenInvalid, err.Error()))
		}
		signature := c.Request().Header.Peek(common.LineSignatureHeader)
		m.ZapLogger.Debug(fmt.Sprintf("%s", signature))
		decoded, err := common.B64ToB(string(signature))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).AuthenticationTokenInvalid, err.Error()))
		}
		secret := os.Getenv("LINE_CHANNEL_SECRET")
		if secret == "" {
			secret = viper.GetString("line.channel.secret")
		}
		hash := hmac.New(sha256.New, []byte(secret))
		_, err = hash.Write(c.Body())
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).AuthenticationTokenInvalid, err.Error()))
		}
		IsValid := hmac.Equal(decoded, hash.Sum(nil))
		if !IsValid {
			err := errors.Wrapf(errors.New(fmt.Sprintf("Token is invalid.")), response.ValidateAuthenticationTokenError)
			return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).AuthenticationTokenInvalid, err.Error()))
		}
		return c.Next()
	}
}
