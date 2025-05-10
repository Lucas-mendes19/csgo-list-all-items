package waxpeer

type ItemsToSellPayload struct {
	Items []ItemToSell `json:"items"`
}

type ItemPricesPayload struct {
	Average int `json:"average"`
	Lowest  int `json:"lowest_price"`
	Currency int `json:"currency"`
}

type ItemToSell struct {
	AssetID      int `json:"item_id"`
	Prices 	   ItemPricesPayload `json:"steam_price"`
}

type OfferRequest struct {
	Offers []Offer `json:"items"`
}

type Offer struct {
	ID       int  `json:"item_id"`
	Price    int `json:"price"`
}
