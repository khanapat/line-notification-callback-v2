package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type PushTextMessageFn func(ctx context.Context, to string, message []TextMessage) error

func NewPushTextMessageFn(lineClient *linebot.Client) PushTextMessageFn {
	return func(ctx context.Context, to string, message []TextMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewTextMessage(value.Text)
			for _, valueEmoji := range value.Emojis {
				message.AddEmoji(linebot.NewEmoji(valueEmoji.Index, valueEmoji.ProductID, valueEmoji.EmojiID))
			}
			messages = append(messages, message)
		}
		_, err := lineClient.PushMessage(to, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type PushStickerMessageFn func(ctx context.Context, to string, message []StickerMessage) error

func NewPushStickerMessageFn(lineClient *linebot.Client) PushStickerMessageFn {
	return func(ctx context.Context, to string, message []StickerMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewStickerMessage(value.PackageID, value.StickerID)
			messages = append(messages, message)
		}
		_, err := lineClient.PushMessage(to, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type PushImageMessageFn func(ctx context.Context, to string, message []ImageMessage) error

func NewPushImageMessageFn(lineClient *linebot.Client) PushImageMessageFn {
	return func(ctx context.Context, to string, message []ImageMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewImageMessage(value.OriginalContentUrl, value.PreviewImageUrl)
			messages = append(messages, message)
		}
		_, err := lineClient.PushMessage(to, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type PushVideoMessageFn func(ctx context.Context, to string, message []VideoMessage) error

func NewPushVideoMessageFn(lineClient *linebot.Client) PushVideoMessageFn {
	return func(ctx context.Context, to string, message []VideoMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewVideoMessage(value.OriginalContentUrl, value.PreviewImageUrl)
			messages = append(messages, message)
		}
		_, err := lineClient.PushMessage(to, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type PushAudioMessageFn func(ctx context.Context, to string, message []AudioMessage) error

func NewPushAudioMessageFn(lineClient *linebot.Client) PushAudioMessageFn {
	return func(ctx context.Context, to string, message []AudioMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewAudioMessage(value.OriginalContentUrl, value.Duration)
			messages = append(messages, message)
		}
		_, err := lineClient.PushMessage(to, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type PushLocationMessageFn func(ctx context.Context, to string, message []LocationMessage) error

func NewPushLocationMessageFn(lineClient *linebot.Client) PushLocationMessageFn {
	return func(ctx context.Context, to string, message []LocationMessage) error {
		var messages []linebot.SendingMessage
		for _, value := range message {
			message := linebot.NewLocationMessage(value.Title, value.Address, value.Latitude, value.Longitude)
			messages = append(messages, message)
		}
		_, err := lineClient.PushMessage(to, messages...).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type PushButtonsTemplateMessageFn func(ctx context.Context, to, altText, thumbnailImageURL, title, text string, message []PostbackActionMessage) error

func NewPushButtonsTemplateMessageFn(lineClient *linebot.Client) PushButtonsTemplateMessageFn {
	return func(ctx context.Context, to, altText, thumbnailImageURL, title, text string, message []PostbackActionMessage) error {
		var actions []linebot.TemplateAction
		for _, value := range message {
			actionByte, err := value.MarshalJSON()
			if err != nil {
				return err
			}
			fmt.Println(string(actionByte))
			var postback PostbackActionMessage
			if err := json.Unmarshal(actionByte, &postback); err != nil {
				return err
			}
			action := linebot.NewPostbackAction(postback.Label, postback.Data, "", postback.DisplayText)
			actions = append(actions, action)
		}
		messages := linebot.NewTemplateMessage(altText, linebot.NewButtonsTemplate(thumbnailImageURL, title, text, actions...))
		_, err := lineClient.PushMessage(to, messages).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

type PushConfirmTemplateMessageFn func(ctx context.Context, to, altText, text string, message []MessageActionMessage) error

func NewPushConfirmTemplateMessageFn(lineClient *linebot.Client) PushConfirmTemplateMessageFn {
	return func(ctx context.Context, to, altText, text string, message []MessageActionMessage) error {
		confirm := linebot.NewConfirmTemplate(text, linebot.NewMessageAction(message[0].Label, message[0].Text), linebot.NewMessageAction(message[1].Label, message[1].Text))
		messages := linebot.NewTemplateMessage(altText, confirm)
		_, err := lineClient.PushMessage(to, messages).Do()
		if err != nil {
			return err
		}
		return nil
	}
}

// type PushFlexMessageFn func(ctx context.Context, to string) error

// func NewPushFlexMessageFn(lineClient *linebot.Client) PushFlexMessageFn {
// 	return func(ctx context.Context, to string) error {
// 		return nil
// 	}
// }
