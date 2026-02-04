package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type INotification interface {
	StartStream(text string) error
	sendMessage(text string) error
}

type Notification struct {
	bot    *bot.Bot
	ctx    context.Context
	update *models.Update
}

func NewNotification() INotification {
	return &Notification{}
}

func (n *Notification) StartStream(text string) error {
	if err := n.sendMessage(text); err != nil {
		return err
	}

	return nil
}

func (n *Notification) sendMessage(text string) error {
	_, err := n.bot.SendMessage(n.ctx, &bot.SendMessageParams{
		ChatID: n.update.Message.Chat.ID,
		Text:   n.update.Message.Text,
	})
	if err != nil {
		return err
	}

	return nil
}
