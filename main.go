package main

import (
	"fmt"
	"line-notification/internal/handler"
	"line-notification/logz"
	"line-notification/middleware"
	"line-notification/notification"
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

	middle := middleware.NewMiddleware(logger)

	line := app.Group(viper.GetString("app.context"))

	line.Use(middle.JSONMiddleware())
	line.Use(middle.ContextLocaleMiddleware())
	line.Use(middle.LoggingMiddleware())

	line.Get("/notification/message", handler.Helper(notification.NewNotificationHandler().PushMessage, logger))

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
