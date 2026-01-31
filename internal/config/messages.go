package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Button struct {
	Text         string `yaml:"text"`
	CallbackData string `yaml:"callback_data"`
}

type MessageConfig struct {
	Name    string     `yaml:"name"`
	Text    string     `yaml:"text"`
	Buttons [][]Button `yaml:"buttons"`
}

type Config struct {
	Messages []MessageConfig `yaml:"messages"`
}

var GlobalConfig Config

func LoadMessages(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &GlobalConfig)
}

func GetMessage(name string, vars map[string]string) (*MessageConfig, error) {
	for _, msg := range GlobalConfig.Messages {
		if msg.Name == name {
			// Clone to avoid modifying global config
			newMsg := msg
			newMsg.Text = ReplaceVars(msg.Text, vars)
			
			newButtons := make([][]Button, len(msg.Buttons))
			for i, row := range msg.Buttons {
				newButtons[i] = make([]Button, len(row))
				for j, btn := range row {
					newButtons[i][j] = Button{
						Text:         ReplaceVars(btn.Text, vars),
						CallbackData: ReplaceVars(btn.CallbackData, vars),
					}
				}
			}
			newMsg.Buttons = newButtons
			return &newMsg, nil
		}
	}
	return nil, fmt.Errorf("message %s not found", name)
}

func ReplaceVars(input string, vars map[string]string) string {
	output := input
	for k, v := range vars {
		placeholder := fmt.Sprintf("{%s}", k)
		output = strings.ReplaceAll(output, placeholder, v)
	}
	return output
}
