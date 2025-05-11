package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	MarketCSGO struct {
		APIKey string `json:"apiKey"`
	} `json:"marketcsgo"`
	Shadowpay struct {
		APIKey string `json:"apiKey"`
	} `json:"shadowpay"`
	Waxpeer struct {
		APIKey string `json:"apiKey"`
	} `json:"waxpeer"`
}

func Load() (*Config, error) {
	var config *Config

	file, err := os.Open("db/config.json")
	if err != nil {
		return config, fmt.Errorf("erro ao abrir arquivo de configuração: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	if config.MarketCSGO.APIKey == "" {
		return fmt.Errorf("API key do MarketCSGO não configurada")
	}
	if config.Shadowpay.APIKey == "" {
		return fmt.Errorf("API key do Shadowpay não configurada")
	}
	if config.Waxpeer.APIKey == "" {
		return fmt.Errorf("API key do Waxpeer não configurada")
	}
	return nil
}