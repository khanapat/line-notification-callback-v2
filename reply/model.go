package reply

import "github.com/line/line-bot-sdk-go/v7/linebot"

// Callback
type CallbackNotificationRequest struct {
	Destination string          `json:"destination" example:""`
	Events      []linebot.Event `json:"events"`
}

// GetProfile
// type GetProfileClientResponse struct {
// 	UserID        string `json:"userId" example:"U7f23c5963e6ef29e206e23d7b785660f"`
// 	DisplayName   string `json:"displayName" example:"trust"`
// 	PictureUrl    string `json:"pictureUrl" example:"https://sprofile.line-scdn.net/0hbmBwarO-PUJvDhUZIYVDPR9ePihMf2RQS20gJFhaYSJWN38VEGh1dgoMNHIFanwXQGx1J1lbNHZjHUokcVjBdmg-Y3VWOnwWQ2t1og"`
// 	StatusMessage string `json:"statusMessage" example:"Passion without discipline is just illusion"`
// 	Language      string `json:"language" example:"en"`
// }
