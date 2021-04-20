package main

import (
	"fmt"
	"line-notification/internal/handler"
	"line-notification/internal/line"
	"line-notification/logz"
	"line-notification/middleware"
	"line-notification/notification"
	"line-notification/reply"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
	_ "time/tzdata"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	runtime.GOMAXPROCS(1)
	initTimezone()
	initViper()
}

func main() {
	timeout := viper.GetDuration("app.timeout")

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		Immutable:     true,
		ReadTimeout:   timeout,
		WriteTimeout:  timeout,
		IdleTimeout:   timeout,
	})

	logger, err := logz.NewLogConfig()
	if err != nil {
		log.Fatal(err)
	}

	lineClient, err := line.NewLineConn()
	if err != nil {
		logger.Error(err.Error())
	}

	middle := middleware.NewMiddleware(logger)

	lineApp := app.Group(viper.GetString("app.context"))

	lineApp.Use(middle.JSONMiddleware())
	lineApp.Use(middle.ContextLocaleMiddleware())
	lineApp.Use(middle.LoggingMiddleware())

	callback := lineApp.Group("/")

	callback.Use(middle.LineAuthenticationMiddleware())

	notiHandler := notification.NewNotificationHandler(
		notification.NewPushTextMessageFn(lineClient),
		notification.NewPushStickerMessageFn(lineClient),
		notification.NewPushImageMessageFn(lineClient),
		notification.NewPushVideoMessageFn(lineClient),
		notification.NewPushAudioMessageFn(lineClient),
		notification.NewPushLocationMessageFn(lineClient),
		notification.NewPushButtonsTemplateMessageFn(lineClient),
		notification.NewPushConfirmTemplateMessageFn(lineClient),
	)

	replyHandler := reply.NewReplyhandler(
		reply.NewGetProfileClientFn(lineClient),
		reply.NewReplyTextMessageFn(lineClient),
		reply.NewReplyStickerMessageFn(lineClient),
	)

	lineApp.Post("/notification/text", handler.Helper(notiHandler.TextNotification, logger))
	lineApp.Post("/notification/sticker", handler.Helper(notiHandler.StickerNotification, logger))
	lineApp.Post("/notification/image", handler.Helper(notiHandler.ImageNotification, logger))
	lineApp.Post("/notification/video", handler.Helper(notiHandler.VideoNotification, logger))
	lineApp.Post("/notification/audio", handler.Helper(notiHandler.AudioNotification, logger))
	lineApp.Post("/notification/location", handler.Helper(notiHandler.LocationNotification, logger))
	lineApp.Post("/notification/template/buttons", handler.Helper(notiHandler.ButtonsTemplateNotification, logger))
	lineApp.Post("/notification/template/confirm", handler.Helper(notiHandler.ConfirmTemplateNotification, logger))

	callback.Post("/callback", handler.Helper(replyHandler.CallbackReply, logger))

	logger.Info(fmt.Sprintf("⇨ http server started on [::]:%s", viper.GetString("app.port")))

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", viper.GetString("app.port"))); err != nil {
			logger.Info(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-c:
		logger.Info("terminating: by signal")
	}

	app.Shutdown()

	logger.Info("shutting down")
	os.Exit(0)
}

func initViper() {
	viper.SetDefault("app.name", "line-notification")
	viper.SetDefault("app.port", "9090")
	viper.SetDefault("app.timeout", "60s")
	viper.SetDefault("app.context", "/ktb/blockchain/v1/line")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.env", "dev")

	viper.SetDefault("line.channel.secret", "708b65656b5f8d0ddbedc24db6e483fe")
	viper.SetDefault("line.channel.access-token", "cdsyUJnmmWbfU8zHbWb4pZVjIw8jrMIXOceX0zBP8e/keH6KAnt4TyG6ZCXGstYP03m68Q1BZOU/7/DvmTyKSDR4EnxLxyAuVq94zQjT6HrjI0bMf9XW5spws2hwmD2ebD+OGHKrVVR3u9i2VlrXCAdB04t89/1O/w1cDnyilFU=")
	// viper.SetDefault("line.user-id", "U7f23c5963e6ef29e206e23d7b785660f")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("error loading location 'Asia/Bangkok': %v\n", err)
	}
	time.Local = ict
}
