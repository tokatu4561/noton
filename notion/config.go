package notion

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
    secretToken string
    pageId string
}

const (
	section = "notion"
	iniFilePath = "/.config/noton.ini"
)

// load config file
func loadConfig() (*Config, error) {
	Config := &Config{}

	// Config ini ファイル読み込み
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %v", err)
	}
	iniFileFullPath := homeDir + iniFilePath
	cfg, err := ini.Load(iniFileFullPath)

	// エラー処理
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	secret := cfg.Section(section).Key("secret").String()
	pageId := cfg.Section(section).Key("page").String()

	Config.secretToken = secret
	Config.pageId = pageId
	
	return Config, nil
}