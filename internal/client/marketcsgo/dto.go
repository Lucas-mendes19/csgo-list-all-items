package marketcsgo

type ItemsToSellPayload struct {
	Items []ItemToSell `json:"items"`
}

type ItemToSell struct {
	AssetID      string `json:"id"`
}

type OfferRequest struct {
	Offers []Offer `json:"items"`
}

type Offer struct {
	ID       string  `json:"asset"`
	Price    float32 `json:"price"`
}
