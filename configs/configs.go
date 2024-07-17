package configs

import (
	"fmt"
	cfgUtils "github.com/nekizz/telegram-bot/internal/pkg/cfg"
	"os"
)

type Configuration struct {
	Stage            string `env:"UP_STAGE"`
	Host             string `env:"HOST"`
	Port             int    `env:"PORT"`
	ReadTimeout      int    `env:"READ_TIMEOUT"`
	WriteTimeout     int    `env:"WRITE_TIMEOUT"`
	Debug            bool   `env:"DEBUG"`
	DbLog            bool   `env:"DB_LOG"`
	DbPsn            string `env:"DB_PSN"`
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
	TelegramChatID   int64  `env:"TELEGRAM_CHAT_ID"`
}

// Load returns Configuration struct
func Load() (*Configuration, error) {
	stage := os.Getenv("UP_STAGE")
	if state := os.Getenv("CONFIG_STAGE"); state != "" {
		stage = state
	}

	cfg := new(Configuration)
	if err := cfgUtils.LoadWithAPS(cfg, stage); err != nil {
		return nil, fmt.Errorf("error parsing environment config: %s", err)
	}

	return cfg, nil
}
