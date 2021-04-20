package line

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func NewLineConn() (*linebot.Client, error) {
	secret := os.Getenv("LINE_CHANNEL_SECRET")
	if secret == "" {
		secret = viper.GetString("line.channel.secret")
	}
	accessToken := os.Getenv("LINE_CHANNEL_TOKEN")
	if accessToken == "" {
		accessToken = viper.GetString("line.channel.access-token")
	}
	bot, err := linebot.New(secret, accessToken)
	if err != nil {
		return nil, err
	}
	return bot, nil
}
