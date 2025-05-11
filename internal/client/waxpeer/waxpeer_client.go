package waxpeer

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

func FetchItemsToSell(key string) error {
	query := url.Values{}
	query.Add("api", key)
	query.Add("game", "csgo")

	request := gorequest.New()
	resp, body, errs := request.Get("https://api.waxpeer.com/v1/fetch-my-inventory"+"?"+query.Encode()).
		Set("Accept", "application/json").
		Set("X-Requested-With", "XMLHttpRequest").
		Set("Authorization", "Bearer "+ key).
		EndBytes()

	if len(errs) > 0 {
		return errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid status code " + strconv.Itoa(resp.StatusCode))
	}

	var responseDTO ResponseDTO
		if err := json.Unmarshal(body, &responseDTO); err != nil {
		return err
	}

	if responseDTO.Success == false {
		return errors.New("failed to fetch items: " + responseDTO.Message)
	}

	return nil
}


func ListItemsToSell(key string) (*ItemsToSellPayload, error) {
	query := url.Values{}
	query.Add("api", key)
	query.Add("skip", strconv.Itoa(0))
	query.Add("game", "csgo")

	request := gorequest.New()
	resp, body, errs := request.Get("https://api.waxpeer.com/v1/get-my-inventory"+"?"+query.Encode()).
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

	if responseDTO.Success == false {
		return nil, errors.New("failed to fetch items")
	}

	return &responseDTO, nil
}

func SellItem(logger *logrus.Logger, key string, items ItemsToSellPayload) error {
	urlBasic := "https://api.waxpeer.com/v1/list-items-steam"

	query := url.Values{}
	query.Add("api", key)
	query.Add("game", "csgo")

	var offers []Offer
	for _, item := range items.Items {
		price := item.Prices.Lowest * 3

		offers = append(offers, Offer{
			ID:       item.AssetID,
			Price:    price,
		})
	}

	payload := OfferRequest{Offers: offers}
		
	request := gorequest.New()
	resp, body, errs := request.Post(urlBasic+"?"+query.Encode()).
		Set("Accept", "application/json").
		Set("Authorization", "Bearer "+ key).
		Send(payload).
		EndBytes()

	if len(errs) > 0 {
		return errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid status code " + strconv.Itoa(resp.StatusCode))
	}

	var responseDTO ResponseDTO
		if err := json.Unmarshal(body, &responseDTO); err != nil {
		return err
	}

	if responseDTO.Success == false {
		return errors.New("failed to sell items: " + responseDTO.Message)
	}

	return nil
}