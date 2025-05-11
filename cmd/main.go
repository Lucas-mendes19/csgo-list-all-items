package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Lucas-mendes19/csgo-list-all-items/infra/configs"
	"github.com/Lucas-mendes19/csgo-list-all-items/infra/logger"
	"github.com/Lucas-mendes19/csgo-list-all-items/internal/domain/service"
)

type ServiceError struct {
	Service string
	Err     error
}

func main() {
	config, err := configs.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar configuração: %v\n", err)
		os.Exit(1)
	}

	logger, logFile, err := logger.Setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao configurar logger: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	waxpeerService := service.NewWaxpeerService(logger, config.Waxpeer.APIKey)
	marketcsgoService := service.NewMarketcsgoService(logger, config.MarketCSGO.APIKey)
	shadowpayService := service.NewShadowpayService(logger, config.Shadowpay.APIKey)

	errCh := make(chan ServiceError, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go runService(&wg, "marketcsgo", marketcsgoService.SellItems, errCh)
	go runService(&wg, "shadowpay", shadowpayService.SellItems, errCh)
	go runService(&wg, "waxpeer", waxpeerService.SellItems, errCh)

	go func() {
		wg.Wait()
		close(errCh)
	}()

	hasError := false
	for se := range errCh {
		if se.Err != nil {
			logger.Errorf("Erro no serviço %s: %v", se.Service, se.Err)
			hasError = true
		}
	}

	if hasError {
		os.Exit(1)
	}
}

func runService(wg *sync.WaitGroup, serviceName string, fn func() error, errCh chan<- ServiceError) {
	defer wg.Done()
	err := fn()
	errCh <- ServiceError{
		Service: serviceName,
		Err:     err,
	}
}