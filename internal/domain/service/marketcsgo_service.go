package service

import (
	"github.com/Lucas-mendes19/csgo-list-all-items/internal/client/marketcsgo"
	"github.com/sirupsen/logrus"
)

type MarketcsgoService struct {
	logger *logrus.Logger
	key    string
}

func NewMarketcsgoService(logger *logrus.Logger, key string) *MarketcsgoService {
	return &MarketcsgoService{
		logger: logger,
		key:    key,
	}
}

func (s *MarketcsgoService) SellItems() error {
	Items, err := marketcsgo.ListItemsToSell(s.key)
	if err != nil {
		s.logger.Error("MarketCSGO - Error listing items:", err)
		return err
	}

	if len(Items.Items) == 0 {
		s.logger.Info("MarketCSGO - No items to sell")
		return nil
	}

	err = marketcsgo.SellItem(s.logger, s.key, *Items)
	if err != nil {
		s.logger.Error("MarketCSGO - Error selling item:", err)
		return nil
	}

	s.logger.Info("MarketCSGO - Items sold successfully")

	return nil
}