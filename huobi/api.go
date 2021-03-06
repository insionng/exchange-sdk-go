package huobi

import (
	"fmt"
)

// GetAllAccountID 获取用户的所有accountid
// GET /v1/account/accounts 查询当前用户的所有账户(即account-id)，Pro站和HADAX account-id通用
func (c *Client) GetAllAccountID() ([]Account, error) {

	r := struct {
		Status string    `json:"status"`
		Data   []Account `json:"data"`
		Errmsg string    `json:"err-msg"`
	}{}
	e := c.doHTTP("GET", "/v1/account/accounts", nil, &r)
	if e != nil {
		return nil, e
	}
	if r.Status != "ok" {
		return nil, fmt.Errorf(r.Errmsg)
	}
	return r.Data, nil
}

// BalanceInfo 查询指定账户的余额
// @parm ishadax true 从HADAX站查询, false pro 站查询
// @parm accountID GetAllAccountID函数返回的id
func (c *Client) BalanceInfo(ishadax bool, accountID int64) ([]Balance, error) {
	f := func() string {
		if ishadax {
			return "/v1/hadax/account/accounts/"
		}
		return "/v1/account/accounts/"
	}

	r := struct {
		Status string `json:"status"`
		Data   struct {
			List []Balance `json:"list"`
		} `json:"data"`

		Errmsg string `json:"err-msg"`
	}{}

	path := fmt.Sprintf("%s%d/balance", f(), accountID)
	e := c.doHTTP("GET", path, nil, &r)
	if e != nil {
		return nil, e
	}
	if r.Status != "ok" {
		return nil, fmt.Errorf(r.Errmsg)
	}
	return r.Data.List, nil
}

// InsertOrder 下单
// @return string: orderNo
func (c *Client) InsertOrder(ishadax bool, req InsertOrderReq) (string, error) {
	req.Source = "api"
	mapParams := if2map(req)

	f := func() string {
		if ishadax {
			return "/v1/hadax/order/orders/place"
		}
		return "/v1/order/orders/place"
	}
	r := struct {
		Status string `json:"status"`
		Errmsg string `json:"err-msg"`
		Data   string `json:"data"`
	}{}
	e := c.doHTTP("POST", f(), mapParams, &r)
	if e != nil {
		return "", e
	}
	if r.Status != "ok" {
		return "", fmt.Errorf(r.Errmsg)
	}
	return r.Data, nil
}

// CancelOrder 撤销一个订单请求
// 注意，返回OK表示撤单请求成功。订单是否撤销成功请调用订单查询接口查询该订单状态
func (c *Client) CancelOrder(orderno string) error {
	path := fmt.Sprintf("/v1/order/orders/%s/submitcancel", orderno)
	r := struct {
		Status string `json:"status"`
		Errmsg string `json:"err-msg"`
		Data   string `json:"data"`
	}{}
	e := c.doHTTP("POST", path, nil, &r)
	if e != nil {
		return e
	}
	if r.Status != "ok" {
		return fmt.Errorf(r.Errmsg)
	}
	return nil
}

// GetOrderDetail 查询某个订单详情
func (c *Client) GetOrderDetail(orderno string) (OrderDetail, error) {
	path := fmt.Sprintf("/v1/order/orders/%s", orderno)
	r := struct {
		Status string      `json:"status"`
		Errmsg string      `json:"err-msg"`
		Data   OrderDetail `json:"data"`
	}{}
	e := c.doHTTP("GET", path, nil, &r)
	if e != nil {
		return OrderDetail{}, e
	}
	if r.Status != "ok" {
		return OrderDetail{}, fmt.Errorf(r.Errmsg)
	}
	return r.Data, nil
}

// GetMatchDetail 查询某个订单的成交明细
func (c *Client) GetMatchDetail(orderno string) ([]MatchDetail, error) {
	path := fmt.Sprintf("/v1/order/orders/%s/matchresults", orderno)
	r := struct {
		Status string        `json:"status"`
		Errmsg string        `json:"err-msg"`
		Data   []MatchDetail `json:"data"`
	}{}
	e := c.doHTTP("GET", path, nil, &r)
	if e != nil {
		return nil, e
	}
	if r.Status != "ok" {
		return nil, fmt.Errorf(r.Errmsg)
	}
	return r.Data, nil
}

// GetOrders 查询当前委托、历史委托
// @parm symbol 交易对		btcusdt, bchbtc, rcneth ...
// @parm status 查询的订单状态组合，使用','分割
//				pre-submitted 准备提交,
//				submitted 已提交,
//				partial-filled 部分成交,
//				partial-canceled 部分成交撤销,
//				filled 完全成交,
//				canceled 已撤销
func (c *Client) GetOrders(symbol, status string) ([]OrderDetail, error) {
	arg := make(map[string]string)
	arg["symbol"] = symbol
	arg["states"] = status
	r := struct {
		Status string        `json:"status"`
		Errmsg string        `json:"err-msg"`
		Data   []OrderDetail `json:"data"`
	}{}
	e := c.doHTTP("GET", "/v1/order/orders", arg, &r)
	if e != nil {
		return nil, e
	}
	if r.Status != "ok" {
		return nil, fmt.Errorf(r.Errmsg)
	}
	return r.Data, nil
}

// TODO: 全部撤单
// TODO: 查询当前成交、历史成交
