package lib

// MarketOrder contains an order (offer or request)
type MarketOrder struct {
	ID               int    `json:"Id"`
	ItemID           string `json:"ItemTypeId"`
	GroupTypeId      string `json:"ItemGroupTypeId"`
	LocationID       int    `json:"LocationId"`
	QualityLevel     int    `json:"QualityLevel"`
	EnchantmentLevel int    `json:"EnchantmentLevel"`
	Price            int    `json:"UnitPriceSilver"`
	Amount           int    `json:"Amount"`
	AuctionType      string `json:"AuctionType"`
	Expires          string `json:"Expires"`
}

// MarketUpload contains a list of orders
type MarketUpload struct {
	Orders []*MarketOrder `json:"Orders"`
}
