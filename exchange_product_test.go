package binance

import "strconv"
import "testing"
import "context"

func TestExchangeProduct(t *testing.T) {
	// t.Fatal("not implemented")
	//t.Log("test success")
	cc := NewClient("", "")
	product, err := cc.NewListProductService().
		Do(context.Background())

	t.Log(err)
	exchangeinfo, err := cc.NewExchangeInfoService().
		Do(context.Background())

	for _, ex := range exchangeinfo.Symbols {
		for _, pr := range product.Data {

			if ex.Symbol == pr.Symbol {
				//t.Log(ex)
				//t.Log(pr)

				pr_min_trade := pr.MinTrade
				ex_min_trade := 0.0
				pr_tick_size := pr.TickSize
				ex_tick_size := 0.0
				for _, fi := range ex.Filters {
					if fi["filterType"] == "LOT_SIZE" {
						ex_min_trade, _ = strconv.ParseFloat(fi["minQty"], 64)
					} else if fi["filterType"] == "PRICE_FILTER" {

						ex_tick_size, _ = strconv.ParseFloat(fi["tickSize"], 64)
					}
				}
				if pr_min_trade != ex_min_trade {
					t.Log("min trade error", pr.Symbol, pr_min_trade, ex_min_trade)
					continue
				}
				if pr_tick_size != ex_tick_size {
					t.Log("tick size error", pr.Symbol, pr_tick_size, ex_tick_size)
					continue
				}
				t.Log(pr.Symbol, "check ok ", pr.Symbol, pr_min_trade, pr_tick_size)

			}
		}
	}
	t.Log(err)
}
