package huobi

import (
	"fmt"
)

// SubMarketKLine 查询市场K线图
// period 可选 1min, 5min, 15min, 30min, 60min, 1day, 1mon, 1week, 1year
func (c *WSSClient) SubMarketKLine(symbol string, period string) (<-chan []byte, error) {
	cid, conn, err := c.connect()
	if err != nil {
		return nil, err
	}

	topic := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	req := struct {
		Topic string `json:"sub"`
		ID    string `json:"id"`
	}{topic, c.generateClientID()}

	err = conn.WriteJSON(req)
	if err != nil {
		c.Close()
		return nil, err
	}

	result := make(chan []byte)
	go c.start(topic, cid, result)
	return result, nil
}

// SubMarketDepth 查询市场深度数据
// type 可选值：{ step0, step1, step2, step3, step4, step5 } （合并深度0-5）；
// step0时，不合并深度
func (c *WSSClient) SubMarketDepth(symbol string, typ string) (<-chan []byte, error) {
	cid, conn, err := c.connect()
	if err != nil {
		return nil, err
	}

	topic := fmt.Sprintf("market.%s.depth.%s", symbol, typ)
	req := struct {
		Topic string `json:"sub"`
		ID    string `json:"id"`
	}{topic, c.generateClientID()}

	err = conn.WriteJSON(req)
	if err != nil {
		c.Close()
		return nil, err
	}

	result := make(chan []byte)
	go c.start(topic, cid, result)
	return result, nil
}

// SubTradeDetail 查询交易详细数据
func (c *WSSClient) SubTradeDetail(symbol string) (<-chan []byte, error) {
	cid, conn, err := c.connect()
	if err != nil {
		return nil, err
	}

	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	req := struct {
		Topic string `json:"sub"`
		ID    string `json:"id"`
	}{topic, c.generateClientID()}

	err = conn.WriteJSON(req)
	if err != nil {
		c.Close()
		return nil, err
	}

	result := make(chan []byte)
	go c.start(topic, cid, result)
	return result, nil
}

// SubMarketDetail 查询市场详情数据
func (c *WSSClient) SubMarketDetail(symbol string) (<-chan []byte, error) {
	cid, conn, err := c.connect()
	if err != nil {
		return nil, err
	}

	topic := fmt.Sprintf("market.%s.detail", symbol)
	req := struct {
		Topic string `json:"sub"`
		ID    string `json:"id"`
	}{topic, c.generateClientID()}

	err = conn.WriteJSON(req)
	if err != nil {
		c.Close()
		return nil, err
	}

	result := make(chan []byte)
	go c.start(topic, cid, result)
	return result, nil
}
