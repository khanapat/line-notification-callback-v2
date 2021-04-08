package notification

import (
	"line-notification/internal/handler"
	"net/http"
)

type notificationhandler struct {
}

func NewNotificationHandler() *notificationhandler {
	return &notificationhandler{}
}

func (s *notificationhandler) PushMessage(c *handler.Ctx) error {
	return c.Status(http.StatusOK).JSON("bobo")
}
