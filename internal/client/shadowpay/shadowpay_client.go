package shadowpay

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

func ListItemsToSell(key string) (*ItemsToSellPayload, error) {
	request := gorequest.New()
	resp, body, errs := request.Get("https://api.shadowpay.com/api/v2/user/inventory").
		Set("Accept", "application/json").
		Set("X-Requested-With", "XMLHttpRequest").
		Set("Authorization", "Bearer "+ key).
		EndBytes()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code " + strconv.Itoa(resp.StatusCode))
	}

	var responseDTO ItemsToSellPayload
	if err := json.Unmarshal(body, &responseDTO); err != nil {
		return nil, err
	}

	return &responseDTO, nil
}

func SellItem(logger *logrus.Logger, key string, items ItemsToSellPayload) error {
	urlBasic := "https://api.shadowpay.com/api/v2/user/offers"

	var offers []Offer
	for _, item := range items.Items {
		offers = append(offers, Offer{
			ID:       item.AssetID,
			Price:    item.Price,
			Project:  "csgo",
			Currency: "USD",
		})
	}

	payload := OfferRequest{Offers: offers}
		
	request := gorequest.New()
	resp, body, errs := request.Post(urlBasic).
		Set("Accept", "application/json").
		Set("Authorization", "Bearer "+ key).
		Send(payload).
		EndBytes()

	logger.Info("Shadowpay - response: ", string(body))

	if len(errs) > 0 {
		return errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid status code " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}