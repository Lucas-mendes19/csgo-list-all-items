package service

import (
	"github.com/Lucas-mendes19/csgo-list-all-items/internal/client/shadowpay"
	"github.com/sirupsen/logrus"
)

type ShadowpayService struct {
	logger *logrus.Logger
	key	string
}

func NewShadowpayService(logger *logrus.Logger, key string) *ShadowpayService {
	return &ShadowpayService{
		logger: logger,
		key:    key,
	}
}

func (s *ShadowpayService) SellItems() error {
	items, err := shadowpay.ListItemsToSell(s.key)
	if err != nil {
		s.logger.Error("Shadowpay - Error listing items:", err)
	}

	if len(items.Items) == 0 {
		s.logger.Info("Shadowpay - No items to sell")
		return nil
	}

	err = shadowpay.SellItem(s.logger, s.key, *items)
	if err != nil {
		s.logger.Error("Shadowpay - Error selling item:", err)
		return nil
	}

	s.logger.Info("Shadowpay - Items sold successfully")

	return nil
}