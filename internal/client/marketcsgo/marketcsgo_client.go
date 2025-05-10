package marketcsgo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

func ListItemsToSell(key string) (*ItemsToSellPayload, error) {
	query := url.Values{}
	query.Add("key", key)

	request := gorequest.New()
	resp, body, errs := request.Get("https://market.csgo.com/api/v2/my-inventory"+"?"+query.Encode()).
		Set("Accept", "application/json").
		Set("X-Requested-With", "XMLHttpRequest").
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
	urlBasic := "https://market.csgo.com/api/v2/mass-add-to-sale"

	query := url.Values{
		"key":   {key},
		"cur":   {"USD"},
	}

	var offers []Offer
	for _, item := range items.Items {
		offers = append(offers, Offer{
			ID:       item.AssetID,
			Price:    5000000,
		})
	}

	payload := OfferRequest{Offers: offers}
		
	request := gorequest.New()
	resp, body, errs := request.Post(urlBasic+"?"+query.Encode()).
		Set("Accept", "application/json").
		Set("Authorization", "Bearer "+ key).
		Send(payload).
		EndBytes()

	logger.Info("Marketcsgo - response: ", string(body))

	if len(errs) > 0 {
		return errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid status code " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}