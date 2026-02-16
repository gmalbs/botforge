package bot

import (
	"strconv"

	"github.com/gmalbs/botforge/internal/config"
	"github.com/gmalbs/botforge/internal/database"
	"github.com/gmalbs/botforge/internal/database/models"

	"github.com/amarnathcjd/gogram/telegram"
)

type Module interface {
	Register(client *telegram.Client)
}

var Modules []Module

func RegisterModule(m Module) {
	Modules = append(Modules, m)
}

func InitBot(token string) (*telegram.Client, error) {
	client, err := telegram.NewClient(telegram.ClientConfig{
		// AppID and AppHash are required by gogram even for bots
		AppID:   int32(config.AppID),
		AppHash: config.AppHash,
	})
	if err != nil {
		return nil, err
	}

	err = client.LoginBot(token)
	if err != nil {
		return nil, err
	}

	// Register all modules
	for _, m := range Modules {
		m.Register(client)
	}

	return client, nil
}

// Helper to check maintenance and get user
func GetUserAndCheckMaintenance(tgUser telegram.User) (*models.User, bool) {
	uObj := tgUser.(*telegram.UserObj)
	var user models.User
	database.DB.Where("user_id = ?", uObj.ID).FirstOrCreate(&user, models.User{
		UserID:    uObj.ID,
		FirstName: uObj.FirstName,
		Username:  uObj.Username,
	})

	var settings models.BotSettings
	database.DB.First(&settings)

	if settings.MaintenanceMode && !user.IsAdmin {
		return &user, true
	}

	return &user, false
}

func SendConfigMessage(client *telegram.Client, chatID int64, msgName string, vars map[string]string) error {
	msgCfg, err := config.GetMessage(msgName, vars)
	if err != nil {
		return err
	}

	kb := telegram.NewKeyboard()
	for _, row := range msgCfg.Buttons {
		var buttons []telegram.KeyboardButton
		for _, btn := range row {
			buttons = append(buttons, telegram.Button.Data(btn.Text, btn.CallbackData))
		}
		kb.AddRow(buttons...)
	}

	_, err = client.SendMessage(chatID, msgCfg.Text, &telegram.SendOptions{
		ParseMode:   "html",
		ReplyMarkup: kb.Build(),
	})
	return err
}

func EditConfigMessage(client *telegram.Client, chatID int64, messageID int32, msgName string, vars map[string]string) error {
	msgCfg, err := config.GetMessage(msgName, vars)
	if err != nil {
		return err
	}

	kb := telegram.NewKeyboard()
	for _, row := range msgCfg.Buttons {
		var buttons []telegram.KeyboardButton
		for _, btn := range row {
			buttons = append(buttons, telegram.Button.Data(btn.Text, btn.CallbackData))
		}
		kb.AddRow(buttons...)
	}

	_, err = client.EditMessage(chatID, messageID, msgCfg.Text, &telegram.SendOptions{
		ParseMode:   "html",
		ReplyMarkup: kb.Build(),
	})
	return err
}

func GetDefaultVars(u *models.User) map[string]string {
	return map[string]string{
		"firstName": u.FirstName,
		"userID":    strconv.FormatInt(u.UserID, 10),
		"username":  u.Username,
	}
}
