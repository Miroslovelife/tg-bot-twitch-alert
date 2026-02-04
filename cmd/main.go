package main

import (
	"github.com/Miroslovelife/tg-bot-twitch-alert/internal/delivery/http/handler"
	"github.com/Miroslovelife/tg-bot-twitch-alert/internal/delivery/http/server"
)

func main() {
	notifHandler := handler.NewNotificationHandler()

	s := server.NewServer()

}
