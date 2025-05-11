package service

import (
	"github.com/Lucas-mendes19/csgo-list-all-items/internal/client/waxpeer"
	"github.com/sirupsen/logrus"
)

type WaxpeerService struct {
	logger *logrus.Logger
	key    string
}

func NewWaxpeerService(logger *logrus.Logger, key string) *WaxpeerService {
	return &WaxpeerService{
		logger: logger,
		key:    key,
	}
}

func (s *WaxpeerService) SellItems() error {
	s.logger.Info("Waxpeer - Starting to sell items")

	err := waxpeer.FetchItemsToSell(s.key)
	if err != nil {
		s.logger.Error("Waxpeer - Error fetching items:", err)
	}

	items, err := waxpeer.ListItemsToSell(s.key)
	if err != nil {
		s.logger.Error("Waxpeer - Error listing items:", err)
		return err
	}

	if len(items.Items) == 0 {
		s.logger.Info("Waxpeer - No items to sell")
		return nil
	}

	err = waxpeer.SellItem(s.logger, s.key, *items)
	if err != nil {
		s.logger.Error("Waxpeer - Error selling item:", err)
		return nil
	}

	s.logger.Info("Waxpeer - Items sold successfully")
	
	return nil
}