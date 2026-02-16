package modules

import (
	"fmt"

	"github.com/gmalbs/botforge/internal/bot"
	"github.com/gmalbs/botforge/internal/config"
	"github.com/gmalbs/botforge/internal/database"
	"github.com/gmalbs/botforge/internal/database/models"

	"github.com/amarnathcjd/gogram/telegram"
)

type AdminModule struct{}

func init() {
	bot.RegisterModule(&AdminModule{})
}

func (m *AdminModule) Register(client *telegram.Client) {
	client.OnCommand("admin", m.handleAdmin)
	client.OnCommand("maintenance", m.toggleMaintenance)
	client.OnCommand("setadmin", m.setAdmin)
}

func (m *AdminModule) handleAdmin(ctx *telegram.NewMessage) error {
	user, _ := bot.GetUserAndCheckMaintenance(ctx.Sender)
	if !user.IsAdmin && user.UserID != config.OwnerID {
		_, err := ctx.Reply("❌ Você não tem permissão para usar este comando.")
		return err
	}

	var settings models.BotSettings
	database.DB.First(&settings)

	status := "Desativado"
	if settings.MaintenanceMode {
		status = "Ativado"
	}

	text := fmt.Sprintf("🛠 <b>Painel Admin</b>\n\nModo Manutenção: <b>%s</b>\n\nComandos:\n/maintenance - Alternar manutenção\n/setadmin [ID] - Dar admin a um usuário", status)
	_, err := ctx.Reply(text, &telegram.SendOptions{ParseMode: "html"})
	return err
}

func (m *AdminModule) toggleMaintenance(ctx *telegram.NewMessage) error {
	user, _ := bot.GetUserAndCheckMaintenance(ctx.Sender)
	if !user.IsAdmin && user.UserID != config.OwnerID {
		return nil
	}

	var settings models.BotSettings
	database.DB.First(&settings)
	settings.MaintenanceMode = !settings.MaintenanceMode
	database.DB.Save(&settings)

	status := "desativado"
	if settings.MaintenanceMode {
		status = "ativado"
	}

	_, err := ctx.Reply(fmt.Sprintf("✅ Modo manutenção %s!", status))
	return err
}

func (m *AdminModule) setAdmin(ctx *telegram.NewMessage) error {
	user, _ := bot.GetUserAndCheckMaintenance(ctx.Sender)
	if !user.IsAdmin && user.UserID != 0 && user.UserID != config.OwnerID {
		return nil
	}

	args := ctx.ArgsList()
	if len(args) == 0 {
		_, err := ctx.Reply("Uso: /setadmin [ID]")
		return err
	}

	targetID := args[0]
	database.DB.Model(&models.User{}).Where("user_id = ?", targetID).Update("is_admin", true)

	_, err := ctx.Reply(fmt.Sprintf("✅ Usuário %s agora é admin!", targetID))
	return err
}
