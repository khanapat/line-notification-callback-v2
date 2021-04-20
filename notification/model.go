package notification

import (
	"encoding/json"
	"fmt"
	"line-notification/common"
	"line-notification/response"
	"unicode/utf8"

	"github.com/pkg/errors"
)

// Text
type TextNotificationRequest struct {
	To      string        `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	Message []TextMessage `json:"message"`
}

type TextMessage struct {
	Text   string  `json:"text" example:"$ trust $"`
	Emojis []Emoji `json:"emojis"`
}

type Emoji struct {
	Index     int    `json:"index" example:"0"`
	ProductID string `json:"productId" example:"5ac1bfd5040ab15980c9b435"`
	EmojiID   string `json:"emojiId" example:"001"`
}

func (req *TextNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if len(req.Message) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'message' must be REQUIRED field but the input is '%v'.", len(req.Message))), response.ValidateFieldError)
	}
	for index, value := range req.Message {
		if utf8.RuneCountInString(value.Text) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].text' must be REQUIRED field but the input is '%v'.", index, value.Text)), response.ValidateFieldError)
		}
		indexes := common.FindRuneIndex(value.Text, '$')
		if len(indexes) == 0 {
			if len(value.Emojis) != 0 {
				return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].emojis' must be empty array.", index)), response.ValidateFieldError)
			}
		} else {
			if len(indexes) != len(value.Emojis) {
				return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].emojis' must have %d items but the input is '%d'.", index, len(indexes), len(value.Emojis))), response.ValidateFieldError)
			}
			for indexEmojis, valueEmojis := range value.Emojis {
				if valueEmojis.Index != indexes[indexEmojis] {
					return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].emojis[%d].index' must be %d but the input is '%d'.", index, indexEmojis, indexes[indexEmojis], valueEmojis.Index)), response.ValidateFieldError)
				}
				if utf8.RuneCountInString(valueEmojis.ProductID) == 0 {
					return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].emojis[%d].productId' must be REQUIRED field but the input is '%v'.", index, indexEmojis, valueEmojis.ProductID)), response.ValidateFieldError)
				}
				if utf8.RuneCountInString(valueEmojis.EmojiID) == 0 {
					return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].emojis[%d].emojiId' must be REQUIRED field but the input is '%v'.", index, indexEmojis, valueEmojis.EmojiID)), response.ValidateFieldError)
				}
			}
		}
	}
	return nil
}

// Sticker
type StickerNotificationRequest struct {
	To      string           `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	Message []StickerMessage `json:"message"`
}

type StickerMessage struct {
	PackageID string `json:"packageId" example:"446"`
	StickerID string `json:"stickerId" example:"1988"`
}

func (req *StickerNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if len(req.Message) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'message' must be REQUIRED field but the input is '%v'.", len(req.Message))), response.ValidateFieldError)
	}
	for index, value := range req.Message {
		if utf8.RuneCountInString(value.PackageID) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].packageId' must be REQUIRED field but the input is '%v'.", index, value.PackageID)), response.ValidateFieldError)
		}
		if utf8.RuneCountInString(value.StickerID) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].stickerId' must be REQUIRED field but the input is '%v'.", index, value.StickerID)), response.ValidateFieldError)
		}
	}
	return nil
}

// Image
type ImageNotificationRequest struct {
	To      string         `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	Message []ImageMessage `json:"message"`
}

type ImageMessage struct {
	OriginalContentUrl string `json:"originalContentUrl" example:"https://example.com/original.jpg"`
	PreviewImageUrl    string `json:"previewImageUrl" example:"https://example.com/preview.jpg"`
}

func (req *ImageNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if len(req.Message) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'message' must be REQUIRED field but the input is '%v'.", len(req.Message))), response.ValidateFieldError)
	}
	for index, value := range req.Message {
		if utf8.RuneCountInString(value.OriginalContentUrl) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].originalContentUrl' must be REQUIRED field but the input is '%v'.", index, value.OriginalContentUrl)), response.ValidateFieldError)
		}
		if utf8.RuneCountInString(value.PreviewImageUrl) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].previewImageUrl' must be REQUIRED field but the input is '%v'.", index, value.PreviewImageUrl)), response.ValidateFieldError)
		}
	}
	return nil
}

// Video
type VideoNotificationRequest struct {
	To      string         `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	Message []VideoMessage `json:"message"`
}

type VideoMessage struct {
	OriginalContentUrl string `json:"originalContentUrl" example:"https://example.com/original.mp4"`
	PreviewImageUrl    string `json:"previewImageUrl" example:"https://example.com/preview.jpg"`
}

func (req *VideoNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if len(req.Message) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'message' must be REQUIRED field but the input is '%v'.", req.Message)), response.ValidateFieldError)
	}
	for index, value := range req.Message {
		if utf8.RuneCountInString(value.OriginalContentUrl) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].originalContentUrl' must be REQUIRED field but the input is '%v'.", index, value.OriginalContentUrl)), response.ValidateFieldError)
		}
		if utf8.RuneCountInString(value.PreviewImageUrl) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].previewImageUrl' must be REQUIRED field but the input is '%v'.", index, value.PreviewImageUrl)), response.ValidateFieldError)
		}
	}
	return nil
}

// Audio
type AudioNotificationRequest struct {
	To      string         `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	Message []AudioMessage `json:"message"`
}

type AudioMessage struct {
	OriginalContentUrl string `json:"originalContentUrl" example:"https://example.com/original.m4a"`
	Duration           int    `json:"duration" example:"60000"`
}

func (req *AudioNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if len(req.Message) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'message' must be REQUIRED field but the input is '%v'.", req.Message)), response.ValidateFieldError)
	}
	for index, value := range req.Message {
		if utf8.RuneCountInString(value.OriginalContentUrl) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].originalContentUrl' must be REQUIRED field but the input is '%v'.", index, value.OriginalContentUrl)), response.ValidateFieldError)
		}
		if value.Duration < 1 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].duration' must be REQUIRED field but the input is '%v'.", index, value.Duration)), response.ValidateFieldError)
		}
	}
	return nil
}

// Location
type LocationNotificationRequest struct {
	To      string            `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	Message []LocationMessage `json:"message"`
}

type LocationMessage struct {
	Title     string  `json:"title" example:"my location"`
	Address   string  `json:"address" example:"1-6-1 Yotsuya, Shinjuku-ku, Tokyo, 160-0004, Japan"`
	Latitude  float64 `json:"latitude" example:"35.687574"`
	Longitude float64 `json:"longitude" example:"139.72922"`
}

func (req *LocationNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if len(req.Message) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'message' must be REQUIRED field but the input is '%v'.", req.Message)), response.ValidateFieldError)
	}
	for index, value := range req.Message {
		if utf8.RuneCountInString(value.Title) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].title' must be REQUIRED field but the input is '%v'.", index, value.Title)), response.ValidateFieldError)
		}
		if utf8.RuneCountInString(value.Address) == 0 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].address' must be REQUIRED field but the input is '%v'.", index, value.Address)), response.ValidateFieldError)
		}
		if value.Latitude < 1 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].latitude' must be REQUIRED field but the input is '%v'.", index, value.Latitude)), response.ValidateFieldError)
		}
		if value.Longitude < 1 {
			return errors.Wrapf(errors.New(fmt.Sprintf("'message[%d].longitude' must be REQUIRED field but the input is '%v'.", index, value.Longitude)), response.ValidateFieldError)
		}
	}
	return nil
}

// ButtonsTemplate
type ButtonsTemplateNotificationRequest struct {
	To                string                  `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	AltText           string                  `json:"altText" example:"This is a buttons template"`
	ThumbnailImageURL string                  `json:"thumbnailImageURL" example:"https://example.com/bot/images/image.jpg"`
	Title             string                  `json:"title" example:"Menu"`
	Text              string                  `json:"text" example:"Please select"`
	Actions           []PostbackActionMessage `json:"actions"`
}

func (req *ButtonsTemplateNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.AltText) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'altText' must be REQUIRED field but the input is '%v'.", req.AltText)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.ThumbnailImageURL) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'thumbnailImageURL' must be REQUIRED field but the input is '%v'.", req.ThumbnailImageURL)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Title) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'title' must be REQUIRED field but the input is '%v'.", req.Title)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Text) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'text' must be REQUIRED field but the input is '%v'.", req.Text)), response.ValidateFieldError)
	}
	if len(req.Actions) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'actions' must be REQUIRED field but the input is '%v'.", len(req.Actions))), response.ValidateFieldError)
	}
	return nil
}

// ConfirmTemplate
type ConfirmTemplateNotificationRequest struct {
	To      string                 `json:"to" example:"U7f23c5963e6ef29e206e23d7b785660f"`
	AltText string                 `json:"altText" example:"this is a confirm template"`
	Text    string                 `json:"text" example:"Are you sure?"`
	Actions []MessageActionMessage `json:"actions"`
}

func (req *ConfirmTemplateNotificationRequest) validate() error {
	if utf8.RuneCountInString(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'.", req.To)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.AltText) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'altText' must be REQUIRED field but the input is '%v'.", req.AltText)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Text) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'text' must be REQUIRED field but the input is '%v'.", req.Text)), response.ValidateFieldError)
	}
	if len(req.Actions) != 2 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'actions' must be 2 field but the input is '%v'.", len(req.Actions))), response.ValidateFieldError)
	}
	return nil
}

// Action
type ActionMessage interface {
	json.Marshaler
	TemplateAction()
}

type PostbackActionMessage struct {
	Label       string `json:"label" example:"Buy"`
	Data        string `json:"data" example:"action=buy&itemid=111"`
	DisplayText string `json:"displayText" example:"Buy"`
}

func (t *PostbackActionMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Label       string `json:"label" example:"Buy"`
		Data        string `json:"data" example:"action=buy&itemid=111"`
		DisplayText string `json:"displayText" example:"Buy"`
	}{
		Label:       t.Label,
		Data:        t.Data,
		DisplayText: t.DisplayText,
	})
}

func (*PostbackActionMessage) TemplateAction() {}

type MessageActionMessage struct {
	Label string `json:"label" example:"Yes"`
	Text  string `json:"text" example:"Yes"`
}

func (t *MessageActionMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Label string `json:"label" example:"Yes"`
		Text  string `json:"text" example:"Yes"`
	}{
		Label: t.Label,
		Text:  t.Text,
	})
}

func (*MessageActionMessage) TemplateAction() {}
