package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/Lucas-mendes19/csgo-list-all-items/internal/client/marketcsgo"
	"github.com/Lucas-mendes19/csgo-list-all-items/internal/client/shadowpay"
	"github.com/Lucas-mendes19/csgo-list-all-items/internal/client/waxpeer"
	log "github.com/sirupsen/logrus"
)

var (
	logger *log.Logger
	file   *os.File
	config Config
)

type Config struct {
	MarketCSGO struct {
		APIKey string
	} `json:"marketcsgo"`
	Shadowpay struct {
		APIKey string
	} `json:"shadowpay"`
	Waxpeer struct {
		APIKey string
	} `json:"waxpeer"`
}

func main() {
	loadConfig()
	setupLogger()
	
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		sellItemMarketcsgo()
	}()

	go func() {
		defer wg.Done()
		sellItemShadowpay()
	}()

	go func() {
		defer wg.Done()
		sellItemWaxpeer()
	}()

	wg.Wait()
	file.Close()

	file.Close()
}

func loadConfig() (Config, error) {
	file, err := os.Open("db/config.json")
	if err != nil {
		return config, fmt.Errorf("erro ao abrir arquivo de configuração: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("erro ao decodificar JSON: %v", err)
	}

	return config, nil
}

func setupLogger() {
	logger = log.New()

	// Configuração do logger
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("Erro ao abrir o arquivo de log: %v", err)
	}

	logger.SetOutput(file)
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	logger.Info("Aplicação iniciada")
}

func sellItemMarketcsgo() {	
	Items, err := marketcsgo.ListItemsToSell(config.MarketCSGO.APIKey)
	if err != nil {
		logger.Error("MarketCSGO - Error listing items:", err)
		return
	}

	if len(Items.Items) == 0 {
		logger.Info("MarketCSGO - No items to sell")
		return
	}

	err = marketcsgo.SellItem(logger, config.MarketCSGO.APIKey, *Items)
	if err != nil {
		logger.Error("MarketCSGO - Error selling item:", err)
		return
	}

	logger.Info("MarketCSGO - Items sold successfully")
}

func sellItemShadowpay() {
	items, err := shadowpay.ListItemsToSell(config.Shadowpay.APIKey)
	if err != nil {
		logger.Error("Shadowpay - Error listing items:", err)
	}

	if len(items.Items) == 0 {
		logger.Info("Shadowpay - No items to sell")
		return
	}

	err = shadowpay.SellItem(logger, config.Shadowpay.APIKey, *items)
	if err != nil {
		logger.Error("Shadowpay - Error selling item:", err)
		return
	}

	logger.Info("Shadowpay - Items sold successfully")
}

func sellItemWaxpeer() {
	items, err := waxpeer.ListItemsToSell(config.Waxpeer.APIKey)
	if err != nil {
		logger.Error("Waxpeer - Error listing items:", err)
		return
	}

	if len(items.Items) == 0 {
		logger.Info("Waxpeer - No items to sell")
		return
	}

	err = waxpeer.SellItem(logger, config.Waxpeer.APIKey, *items)
	if err != nil {
		logger.Error("Waxpeer - Error selling item:", err)
		return
	}

	logger.Info("Waxpeer - Items sold successfully")
}