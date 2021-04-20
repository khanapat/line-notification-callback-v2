package reply

import (
	"fmt"
	"line-notification/common"
	"line-notification/internal/handler"
	"line-notification/notification"
	"line-notification/response"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type replyhandler struct {
	GetProfileClientFn    GetProfileClientFn
	ReplyTextMessageFn    ReplyTextMessageFn
	ReplyStickerMessageFn ReplyStickerMessageFn
}

func NewReplyhandler(getProfileClientFn GetProfileClientFn, replyTextMessageFn ReplyTextMessageFn, replyStickerMessageFn ReplyStickerMessageFn) *replyhandler {
	return &replyhandler{
		GetProfileClientFn:    getProfileClientFn,
		ReplyTextMessageFn:    replyTextMessageFn,
		ReplyStickerMessageFn: replyStickerMessageFn,
	}
}

func (s *replyhandler) CallbackReply(c *handler.Ctx) error {
	var req CallbackNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).ReplyMessageValidateReq, err.Error()))
	}

	if len(req.Events) != 0 {
		for _, event := range req.Events {
			if event.Type == linebot.EventTypeMessage {
				userProfile, err := s.GetProfileClientFn(c.Context(), event.Source.UserID)
				if err != nil {
					return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).ReplyMessageThirdParty, err.Error()))
				}
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					c.Log().Debug(fmt.Sprintf("Text : %s", message.Text))
					replyText := fmt.Sprintf("Hi! %s $, bobo", userProfile.DisplayName)
					indexs := common.FindRuneIndex(replyText, '$')
					textMessages := []notification.TextMessage{
						{
							Text: replyText,
							Emojis: []notification.Emoji{
								{
									Index:     indexs[0],
									ProductID: "5ac1bfd5040ab15980c9b435",
									EmojiID:   "001",
								},
							},
						},
					}
					if err := s.ReplyTextMessageFn(c.Context(), event.ReplyToken, textMessages); err != nil {
						return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).ReplyMessageThirdParty, err.Error()))
					}
				case *linebot.StickerMessage:
					c.Log().Debug(fmt.Sprintf("StickerID : %s", message.StickerID))
					textMessages := []notification.StickerMessage{
						{
							PackageID: "6136",
							StickerID: "10551380",
						},
					}
					if err := s.ReplyStickerMessageFn(c.Context(), event.ReplyToken, textMessages); err != nil {
						return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).ReplyMessageThirdParty, err.Error()))
					}
				default:
					replyText := fmt.Sprintf("This feature is unavailable.")
					textMessages := []notification.TextMessage{
						{
							Text:   replyText,
							Emojis: []notification.Emoji{},
						},
					}
					if err := s.ReplyTextMessageFn(c.Context(), event.ReplyToken, textMessages); err != nil {
						return c.Status(http.StatusInternalServerError).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).ReplyMessageThirdParty, err.Error()))
					}
				}
			}
		}
	}

	c.Log().Info(fmt.Sprintf("Reply line callback success."))
	return c.Status(http.StatusOK).JSON(response.ResponseContextLocale(c.Context()).ReplyMessageSuccess)
}
