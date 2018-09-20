package binance

import (
	"context"
	"encoding/json"
	"time"
)

// ListBookTickersService list all book tickers
type AssetDetailService struct {
	c *Client
}

var (
	lastCall time.Time = time.Now()
	lastData []byte
)

// Do send request
func (s *AssetDetailService) Do(ctx context.Context, opts ...RequestOption) (asset AssetDetail, err error) {

	var data []byte
	if lastData != nil && time.Since(lastCall).Seconds() < 10.0 {
		data = lastData
	} else {
		r := &request{
			method:   "GET",
			endpoint: "/wapi/v3/assetDetail.html",
			secType:  secTypeSigned,
		}
		data, err = s.c.callAPI(ctx, r, opts...)
		if err != nil {
			return
		}
	}
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &asset)
	if err != nil {
		return
	}
	return
}

/*
{
  "success": true,
  "assetDetail": {
    "CTR": {
      "minWithdrawAmount": "70.00000000",
      "depositStatus": false,
      "withdrawFee": 35,
      "withdrawStatus": true,
      "depositTip": "Delisted, Deposit Suspended"
    },
    "SKY": {
      "minWithdrawAmount": "0.02000000",
      "depositStatus": true,
      "withdrawFee": 0.01,
      "withdrawStatus": true
    }
  }
}
*/
// PriceChangeStats define price change stats
type AssetDetail struct {
	Success     bool             `json:"success"`
	AssetDetail map[string]Asset `json:"assetDetail"`
}

type Asset struct {
	MinWithdrawAmount float64 `json:"minWithdrawAmount"`
	DepositStatus     bool    `json:"depositStatus"`
	WithdrawFee       float64 `json:"withdrawFee"`
	WithdrawStatus    bool    `json:"withdrawStatus"`
}
