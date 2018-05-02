package main

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/blockcdn-go/exchange-sdk-go/binance"
)

func main() {

	pxy, _ := url.Parse("http://127.0.0.1:1080")

	//ctx, _ := context.WithCancel(context.Background())
	// use second return value for cancelling request
	b := binance.NewAPIService(
		nil,
		"https://www.binance.com",
		"X4gV04iX86rr4a0urODBhQrGs0us2MTF5VyELvqE8uGvzM7LtXqg3ckqsevoZKRe",
		"1XlcmtmbGRMmqdDoxeRuzkTYXJZECu1OOF4q1toanmxWFVPjWOJPxPZeDj2jvpRS",
		pxy)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	kech, err := b.TradeWebsocket("BTCUSDT")
	if err != nil {
		panic(err)
	}
	depth, _ := b.DepthWebsocket("BTCUSDT")

	tk, _ := b.TickerWebsocket("BTCUSDT")
	go func() {
		for {
			select {
			case ke := <-kech:
				fmt.Printf("%+v\n", ke)
			case d := <-depth:
				fmt.Printf("%+v\n", d)
			case t := <-tk:
				fmt.Printf("%+v\n", t)
			}
		}
	}()

	fmt.Println("waiting for interrupt")
	<-interrupt
	fmt.Println("canceling context")

	fmt.Println("exit")
	return

	kl, err := b.Klines(binance.KlinesRequest{
		Symbol:   "BNBETH",
		Interval: binance.Hour,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", kl)

	newOrder, err := b.NewOrder(binance.NewOrderRequest{
		Symbol:      "BNBETH",
		Quantity:    1,
		Price:       999,
		Side:        binance.SideSell,
		TimeInForce: binance.GTC,
		Type:        binance.TypeLimit,
		Timestamp:   time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(newOrder)

	res2, err := b.QueryOrder(binance.QueryOrderRequest{
		Symbol:     "BNBETH",
		OrderID:    newOrder.OrderID,
		RecvWindow: 5 * time.Second,
		Timestamp:  time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res2)

	res4, err := b.OpenOrders(binance.OpenOrdersRequest{
		Symbol:     "BNBETH",
		RecvWindow: 5 * time.Second,
		Timestamp:  time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res4)

	res3, err := b.CancelOrder(binance.CancelOrderRequest{
		Symbol:    "BNBETH",
		OrderID:   newOrder.OrderID,
		Timestamp: time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res3)

	res5, err := b.AllOrders(binance.AllOrdersRequest{
		Symbol:     "BNBETH",
		RecvWindow: 5 * time.Second,
		Timestamp:  time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res5[0])

	res6, err := b.Account(binance.AccountRequest{
		RecvWindow: 5 * time.Second,
		Timestamp:  time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res6)

	res7, err := b.MyTrades(binance.MyTradesRequest{
		Symbol:     "BNBETH",
		RecvWindow: 5 * time.Second,
		Timestamp:  time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res7)

	res9, err := b.DepositHistory(binance.HistoryRequest{
		Timestamp:  time.Now(),
		RecvWindow: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res9)

	res8, err := b.WithdrawHistory(binance.HistoryRequest{
		Timestamp:  time.Now(),
		RecvWindow: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res8)

	ds, err := b.StartUserDataStream()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", ds)

	err = b.KeepAliveUserDataStream(ds)
	if err != nil {
		panic(err)
	}

	err = b.CloseUserDataStream(ds)
	if err != nil {
		panic(err)
	}
}
