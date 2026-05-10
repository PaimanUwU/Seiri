package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	NotesDir     string `yaml:"notes_dir"`
	Theme        string `yaml:"theme"`
	ShowIcons    bool   `yaml:"show_icons"`
	SyncInterval int    `yaml:"sync_interval"`
}

func LoadConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(home, ".config", "seiri", "config.yml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return &Config{
			NotesDir:     filepath.Join(home, "Documents", "Notes"),
			Theme:        "catppuccin-mocha",
			ShowIcons:    true,
			SyncInterval: 30,
		}, nil
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	if len(cfg.NotesDir) > 0 && cfg.NotesDir == '~' {
		cfg.NotesDir = filepath.Join(home, cfg.NotesDir[1:])
	}

	return &cfg, nil
}
