package dto

type TradeInput struct {
	OrderID       string  `json:"order_id"`
	InvestorID    string  `json:"investor_id"`
	AssetID       string  `json:"asset_id"`
	OrderType     string  `json:"order_type"`
	CurrentShares int     `json:"current_shares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
}

type OrderOutput struct {
	OrderID           string               `json:"order_id"`
	InvestorID        string               `json:"investor_id"`
	AssetID           string               `json:"asset_id"`
	OrderType         string               `json:"order_type"`
	Status            string               `json:"status"`
	Shares            int                  `json:"shares"`
	Partial           int                  `json:"partial"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerID       string  `json:"buyer_id"`
	SellerID      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}
