package telegram

import (
	"context"
	"os"
	"os/signal"

	telegram "github.com/Miroslovelife/tg-bot-twitch-alert/internal/telegram/handler"
	bot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ITelegramServer interface {
}

type TelegramServer struct {
	Config  *TelegramApiConfig
	Storage *LocalStorage
}

type TelegramApiConfig struct {
	ApiKey   string
	Opts     []bot.Option
	Handlers []telegram.Handler
}

type LocalStorage struct {
	Chats            []*models.ChatFullInfo
	LinkedChatTwitch []*LinkedChatTwitch
}

type LinkedChatTwitch struct {
	Channel    string
	ChatId     int64
	TwitchLink string
}

func NewTelegramServer() ITelegramServer {
	return &TelegramServer{}
}

func (t *TelegramServer) MustInitTelegramServer() (*bot.Bot, context.Context) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(t.Config.ApiKey, t.Config.Opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)

	return b, ctx
}

func (t *TelegramServer) InitHandlers(handlers []telegram.Handler) {
	t.Config.Handlers = handlers

	t.Config.Opts = []bot.Option{
		bot.WithDefaultHandler(t.RegChat),
	}
}

func (t *TelegramServer) RegChat(ctx context.Context, b *bot.Bot, update *models.Update) {
	chat, err := b.GetChat(ctx, &bot.GetChatParams{
		ChatID: update.Message.Chat.ID,
	})
	if err != nil {
	}

	t.Storage.Chats = append(t.Storage.Chats, chat)
}

func (t *TelegramServer) LinkTwitchAccount(ctx context.Context, b *bot.Bot, update *models.Update) {
	chat, err := b.GetChat(ctx, &bot.GetChatParams{
		ChatID: update.Message.Chat.ID,
	})
	if err != nil {
	}

	t.Storage.LinkedChatTwitch = append(t.Storage.LinkedChatTwitch, &LinkedChatTwitch{
		ChatId:     chat.ID,
		TwitchLink: update.Message.Text,
	})
}

func (t *TelegramServer) LinkTelegramChannel(ctx context.Context, b *bot.Bot, update *models.Update) {
	chat, err := b.GetChat(ctx, &bot.GetChatParams{
		ChatID: update.Message.Chat.ID,
	})
	if err != nil {
	}

	exists := map[int64]any{}

	for _, v := range t.Storage.LinkedChatTwitch {
		if _, ok := exists[chat.ID]; ok {
			v.Channel = update.Message.Text

			break
		}

		exists[chat.ID] = ""
	}

}
