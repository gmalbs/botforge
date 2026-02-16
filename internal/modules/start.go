package modules

import (
	"github.com/gmalbs/botforge/internal/bot"

	"github.com/amarnathcjd/gogram/telegram"
)

type StartModule struct{}

func init() {
	bot.RegisterModule(&StartModule{})
}

func (m *StartModule) Register(client *telegram.Client) {
	client.OnCommand("start", m.handleStart)
	client.OnCallback("start", m.handleStartCallback)
	client.OnCallback("help", m.handleHelp)
	client.OnCallback("about", m.handleAbout)
}

func (m *StartModule) handleStart(ctx *telegram.NewMessage) error {
	user, maintenance := bot.GetUserAndCheckMaintenance(ctx.Sender)
	if maintenance {
		return bot.SendConfigMessage(ctx.Client, ctx.ChatID(), "maintenance", nil)
	}

	vars := bot.GetDefaultVars(user)
	return bot.SendConfigMessage(ctx.Client, ctx.ChatID(), "start", vars)
}

func (m *StartModule) handleStartCallback(ctx *telegram.CallbackQuery) error {
	user, maintenance := bot.GetUserAndCheckMaintenance(ctx.Sender)
	if maintenance {
		return bot.EditConfigMessage(ctx.Client, ctx.ChatID, ctx.MessageID, "maintenance", nil)
	}

	vars := bot.GetDefaultVars(user)
	return bot.EditConfigMessage(ctx.Client, ctx.ChatID, ctx.MessageID, "start", vars)
}

func (m *StartModule) handleHelp(ctx *telegram.CallbackQuery) error {
	user, _ := bot.GetUserAndCheckMaintenance(ctx.Sender)
	vars := bot.GetDefaultVars(user)
	return bot.EditConfigMessage(ctx.Client, ctx.ChatID, ctx.MessageID, "help", vars)
}

func (m *StartModule) handleAbout(ctx *telegram.CallbackQuery) error {
	_, err := ctx.Answer("Este é um bot modular feito em Go!", &telegram.CallbackOptions{
		Alert: true,
	})
	return err
}
