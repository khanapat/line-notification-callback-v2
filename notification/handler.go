package notification

import (
	"fmt"
	"line-notification/internal/handler"
	"line-notification/response"
	"net/http"
)

type notificationhandler struct {
	PushTextMessageFn            PushTextMessageFn
	PushStickerMessageFn         PushStickerMessageFn
	PushImageMessageFn           PushImageMessageFn
	PushVideoMessageFn           PushVideoMessageFn
	PushAudioMessageFn           PushAudioMessageFn
	PushLocationMessageFn        PushLocationMessageFn
	PushButtonsTemplateMessageFn PushButtonsTemplateMessageFn
	PushConfirmTemplateMessageFn PushConfirmTemplateMessageFn
}

func NewNotificationHandler(pushTextMessageFn PushTextMessageFn, pushStickerMessageFn PushStickerMessageFn, pushImageMessageFn PushImageMessageFn, pushVideoMessageFn PushVideoMessageFn, pushAudioMessageFn PushAudioMessageFn, pushLocationMessageFn PushLocationMessageFn, pushButtonsTemplateMessageFn PushButtonsTemplateMessageFn, pushConfirmTemplateMessageFn PushConfirmTemplateMessageFn) *notificationhandler {
	return &notificationhandler{
		PushTextMessageFn:            pushTextMessageFn,
		PushStickerMessageFn:         pushStickerMessageFn,
		PushImageMessageFn:           pushImageMessageFn,
		PushVideoMessageFn:           pushVideoMessageFn,
		PushAudioMessageFn:           pushAudioMessageFn,
		PushLocationMessageFn:        pushLocationMessageFn,
		PushButtonsTemplateMessageFn: pushButtonsTemplateMessageFn,
		PushConfirmTemplateMessageFn: pushConfirmTemplateMessageFn,
	}
}

func (s *notificationhandler) TextNotification(c *handler.Ctx) error {
	var req TextNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushTextMessageFn(c.Context(), req.To, req.Message); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}

func (s *notificationhandler) StickerNotification(c *handler.Ctx) error {
	var req StickerNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushStickerMessageFn(c.Context(), req.To, req.Message); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}

func (s *notificationhandler) ImageNotification(c *handler.Ctx) error {
	var req ImageNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushImageMessageFn(c.Context(), req.To, req.Message); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}

func (s *notificationhandler) VideoNotification(c *handler.Ctx) error {
	var req VideoNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushVideoMessageFn(c.Context(), req.To, req.Message); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}

func (s *notificationhandler) AudioNotification(c *handler.Ctx) error {
	var req AudioNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushAudioMessageFn(c.Context(), req.To, req.Message); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}

func (s *notificationhandler) LocationNotification(c *handler.Ctx) error {
	var req LocationNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushLocationMessageFn(c.Context(), req.To, req.Message); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}

func (s *notificationhandler) ButtonsTemplateNotification(c *handler.Ctx) error {
	var req ButtonsTemplateNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushButtonsTemplateMessageFn(c.Context(), req.To, req.AltText, req.ThumbnailImageURL, req.Title, req.Text, req.Actions); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}

func (s *notificationhandler) ConfirmTemplateNotification(c *handler.Ctx) error {
	var req ConfirmTemplateNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageValidateReq, err.Error()))
	}

	if err := s.PushConfirmTemplateMessageFn(c.Context(), req.To, req.AltText, req.Text, req.Actions); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).PushMessageThirdParty, err.Error()))
	}
	c.Log().Info(fmt.Sprintf("Send line notification success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).PushMessageSuccess)
}
