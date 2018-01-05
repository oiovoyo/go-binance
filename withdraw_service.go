package binance

import (
    "context"
    "encoding/json"
    "errors"
)

// CreateWithdrawService create withdraw
type CreateWithdrawService struct {
    c          *Client
    asset      string
    address    string
    amount     string
    name       *string
    addressTag *string
}

// Asset set asset
func (s *CreateWithdrawService) Asset(asset string) *CreateWithdrawService {
    s.asset = asset
    return s
}

// Address set address
func (s *CreateWithdrawService) Address(address string) *CreateWithdrawService {
    s.address = address
    return s
}

// Amount set amount
func (s *CreateWithdrawService) Amount(amount string) *CreateWithdrawService {
    s.amount = amount
    return s
}

// Name set name
func (s *CreateWithdrawService) Name(name string) *CreateWithdrawService {
    s.name = &name
    return s
}

func (s *CreateWithdrawService) AddressTag(addressTag string) *CreateWithdrawService {
    s.addressTag = &addressTag
    return s
}

// Do send request
func (s *CreateWithdrawService) Do(ctx context.Context) (id string, err error) {
    r := &request{
        method:   "POST",
        endpoint: "/wapi/v3/withdraw.html",
        secType:  secTypeSigned,
    }
    m := params{
        "asset":   s.asset,
        "address": s.address,
        "amount":  s.amount,
    }
    if s.name != nil {
        m["name"] = *s.name
    }
    if s.addressTag != nil {
        m["addressTag"] = *s.addressTag
    }
    //r.SetFormParams(m)
    r.SetParams(m)
    data, err := s.c.callAPI(ctx, r)

    if err != nil {
        return
    }
    //fmt.Printf("%s\n",string(data))
    withdraw := new(WithdrawResponse)
    err = json.Unmarshal(data, withdraw)
    if err != nil {
        return
    }

    if withdraw.Success == false {
        err = errors.New(withdraw.Msg)
    }

    return withdraw.Id, nil
}

type WithdrawResponse struct {
    Success bool   `json:"success"`
    Msg     string `json:"msg"`
    Id      string `json:"id"`
}

// ListWithdrawsService list withdraws
type ListWithdrawsService struct {
    c         *Client
    asset     *string
    status    *int
    startTime *int64
    endTime   *int64
}

// Asset set asset
func (s *ListWithdrawsService) Asset(asset string) *ListWithdrawsService {
    s.asset = &asset
    return s
}

// Status set status
func (s *ListWithdrawsService) Status(status int) *ListWithdrawsService {
    s.status = &status
    return s
}

// StartTime set startTime
func (s *ListWithdrawsService) StartTime(startTime int64) *ListWithdrawsService {
    s.startTime = &startTime
    return s
}

// EndTime set endTime
func (s *ListWithdrawsService) EndTime(endTime int64) *ListWithdrawsService {
    s.endTime = &endTime
    return s
}

// Do send request
func (s *ListWithdrawsService) Do(ctx context.Context) (withdraws []*Withdraw, err error) {
    r := &request{
        method:   "GET",
        endpoint: "/wapi/v3/withdrawHistory.html",
        secType:  secTypeSigned,
    }
    if s.asset != nil {
        r.SetParam("asset", *s.asset)
    }
    if s.status != nil {
        r.SetParam("status", *s.status)
    }
    if s.startTime != nil {
        r.SetParam("startTime", *s.startTime)
    }
    if s.endTime != nil {
        r.SetParam("endTime", *s.endTime)
    }
    data, err := s.c.callAPI(ctx, r)
    if err != nil {
        return
    }
    res := new(WithdrawHistoryResponse)
    err = json.Unmarshal(data, res)
    if err != nil {
        return
    }
    return res.Withdraws, nil
}

// WithdrawHistoryResponse define withdraw history response
type WithdrawHistoryResponse struct {
    Withdraws []*Withdraw `json:"withdrawList"`
    Success   bool        `json:"success"`
}

// Withdraw define withdraw info
type Withdraw struct {
    Id         string  `json:"id"`
    Amount     float64 `json:"amount"`
    Address    string  `json:"address"`
    Asset      string  `json:"asset"`
    TxID       string  `json:"txId"`
    ApplyTime  int64   `json:"applyTime"`
    Status     int     `json:"status"`
    AddressTag *string `json:"addressTag"`
}
