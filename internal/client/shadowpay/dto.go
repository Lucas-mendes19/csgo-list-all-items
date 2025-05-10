package shadowpay

type ItemsToSellPayload struct {
	Items []ItemToSell `json:"data"`
}

type ItemToSell struct {
	AssetID      string `json:"asset_id"`
	Price 	     float32 `json:"max_price"`
}

type OfferRequest struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	ID       string  `json:"id"`
	Price    float32 `json:"price"`
	Project  string  `json:"project"`
	Currency string  `json:"currency"`
}
