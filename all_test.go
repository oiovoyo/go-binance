package binance

import "testing"
import "time"
import "context"
import "net/http"
import "net"
import "net/url"

var (
	apiKey    = ""
	secretKey = ""

	proxyString = "https://127.0.0.1:1087"
	proxyUrl, _ = url.Parse(proxyString)
	t           = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 60 * time.Second,
		Proxy:               http.ProxyURL(proxyUrl),
	}
	c0 = &http.Client{
		Transport: t,
		Timeout:   15 * time.Second,
	}
	client = NewClientCustomHttp(apiKey, secretKey, c0)
)

func TestClient_NewListDepositsService(t *testing.T) {
	v, err := client.NewListDepositsService().
		Asset("USDT").
		StartTime(time.Now().AddDate(0, 0, -10).UTC().UnixNano() / int64(time.Millisecond)).
		EndTime(time.Now().UTC().UnixNano() / int64(time.Millisecond)).
		Do(context.Background())
	t.Log(v[0], err)
	t.Log(time.Now().AddDate(0, 0, -1).UTC().UnixNano() / int64(time.Millisecond))
}

func TestClient_NewCreateOrderServiceBuy(t *testing.T) {
	v, err := client.NewCreateOrderService().Symbol("BTCUSDT").
		Side(SideTypeBuy).Type(OrderTypeLimit).
		TimeInForce(TimeInForceGTC).Quantity("0.01").
		Price("5600.009").Do(context.Background())
	t.Log(v, err)
}

func TestClient_NewCreateOrderServiceSell(t *testing.T) {
	v, err := client.NewCreateOrderService().Symbol("BTCUSDT").
		Side(SideTypeSell).Type(OrderTypeLimit).
		TimeInForce(TimeInForceGTC).Quantity("0.00999").
		Price("5900.0").Do(context.Background())
	t.Log(v, err)
}

func TestClient_NewListOpenOrdersService(t *testing.T) {
	v, err := client.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		t.Log(v, err)
		return
	}
	for _, o := range v {
		t.Log("%+v",o)
	}
}

func TestClient_NewCancelOrderService(t *testing.T) {
	v, err := client.NewCancelOrderService().Symbol("BTCUSDT").
	OrigClientOrderID("JwYO8kowMO8SAeZZQ07lvX").
	Do(context.Background())

	t.Log(v, err)
}

func TestClient_NewGetAccountService(t *testing.T) {
	v, err := client.NewGetAccountService().
		Do(context.Background())

	t.Log(v, err)
}


func TestClient_NewListProductService(t *testing.T) {
	v, err := client.NewListProductService().
		Do(context.Background())

	t.Log(v, err)
}

func TestClient_NewNewAccoutStatusService(t *testing.T) {
	 err := client.NewAccoutStatusService().
		Do(context.Background())

	t.Log( err)
}

func TestClient_NewDepositAddressService(t *testing.T) {
	a, err := client.NewDepositAddressService().Asset("ARN").
		Do(context.Background())

	t.Log(a, err)
}