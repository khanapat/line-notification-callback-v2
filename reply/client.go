package reply

import (
	"context"
	"line-notification/notification"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type GetProfileClientFn func(ctx context.Context, userId string) (*linebot.UserProfileResponse, error)

func NewGetProfileClientFn(lineClient *linebot.Client) GetProfileClientFn {
	return func(ctx context.Context, userId string) (*linebot.UserProfileResponse, error) {
		userProfile, err := lineClient.GetProfile(userId).Do()
		if err != nil {
			return nil, err
		}
		return userProfile, nil
	}
}

type ReplyTextMessageFn func(ctx context.Context, replyToken string, message []notification.TextMessage) error

func NewReplyTextMessageFn(lineClient *linebot.Client) ReplyTextMessageFn {
	return func(ctx context.Context, replyToken string, message []notification.TextMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewTextMessage(value.Text)
			for _, valueEmoji := range value.Emojis {
				message.AddEmoji(linebot.NewEmoji(valueEmoji.Index, valueEmoji.ProductID, valueEmoji.EmojiID))
			}
			messages = append(messages, message)
		}
		_, err := lineClient.ReplyMessage(replyToken, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type ReplyStickerMessageFn func(ctx context.Context, replyToken string, message []notification.StickerMessage) error

func NewReplyStickerMessageFn(lineClient *linebot.Client) ReplyStickerMessageFn {
	return func(ctx context.Context, replyToken string, message []notification.StickerMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewStickerMessage(value.PackageID, value.StickerID)
			messages = append(messages, message)
		}
		_, err := lineClient.ReplyMessage(replyToken, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}
