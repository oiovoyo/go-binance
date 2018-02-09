package binance

import (
	"context"
	"encoding/json"
)

// ListBookTickersService list all book tickers
type ListProductService struct {
	c *Client
}

// Do send request
func (s *ListProductService) Do(ctx context.Context, opts ...RequestOption) (p ProductList, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/exchange/public/product",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return
	}
	return
}

/*
{
      "symbol": "BNBBTC",
      "tradedMoney": 537.90333759,
      "baseAssetUnit": "",
      "active": true,
      "minTrade": "1.00000000",
      "baseAsset": "BNB",
      "activeSell": 2479809,
      "withdrawFee": "0",
      "tickSize": "0.00000001",
      "prevClose": 2.0002E-4,
      "activeBuy": 0,
      "volume": "2479809.00000000",
      "high": "0.00022969",
      "lastAggTradeId": 1451985,
      "decimalPlaces": 8,
      "low": "0.00019964",
      "quoteAssetUnit": "฿",
      "matchingUnitType": "STANDARD",
      "close": "0.00020495",
      "quoteAsset": "BTC",
      "open": "0.00020002",
      "status": "TRADING",
      "minQty": "1E-8"
    },

*/
// PriceChangeStats define price change stats
type ProductList struct {
	Data []Product `json:"data"`
}

type Product struct {
	Symbol           string  `json:"symbol"`
	TradedMoney      float64 `json:"tradedMoney"`
	Active           bool    `json:"active"`
	MinTrade         float64 `json:"minTrade,string"`
	BaseAsset        string  `json:"baseAsset"`
	ActiveSell       float64 `json:"activeSell"`
	WithdrawFee      float64 `json:"withdrawFee,string"`
	TickSize         float64 `json:"tickSize,string"`
	PrevClose        float64 `json:"prevClose"`
	ActiveBuy        float64 `json:"activeBuy"`
	Volume           float64 `json:"volume,string"`
	High             float64 `json:"high,string"`
	LastAggTradeId   int64   `json:"lastAggTradeId"`
	DecimalPlaces    int64   `json:"decimalPlaces"`
	QuoteAssetUnit   string  `json:"quoteAssetUnit"`
	MatchingUnitType string  `json:"matchingUnitType"`
	Close            float64 `json:"close,string"`
	QuoteAsset       string  `json:"quoteAsset"`
	Open             float64 `json:"open,string"`
	Status           string  `json:"status"`
	MinQty           float64 `json:"minQty,string"`
}
