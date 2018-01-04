package binance

import (
    "context"
    //"encoding/json"
    //"errors"
    "fmt"
)

// CreateWithdrawService create withdraw
type CreateAccoutStatusService struct {
    c          *Client
}

// Do send request
func (s *CreateAccoutStatusService) Do(ctx context.Context) (err error) {
    r := &request{
        method:   "GET",
        endpoint: "/wapi/v3/accountStatus.html",
        secType:  secTypeSigned,
    }

    //r.SetFormParams(m)
    //r.SetParams(m)
    data, err := s.c.callAPI(ctx, r)

    if err != nil {
        return
    }
    fmt.Printf("%s\n",string(data))


    return
}
