package model

import "time"

type Cart struct {
	ProductIdx int       `json:"itemidx"`
	Quantity   int       `json:"quantity"`
	Addtime    time.Time `json:"addtime"`
}

type CartPrdt struct {
	Cart
	Product
}

type CartDBHandler interface { // DBHandler2 인터페이스에 GetTests, Addtodo 인자에 sessionId를 추가해주어야 한다.
	GetCarts(sessionId string) []*CartPrdt
	AddCart(itemidx int, quantity int, sessionId string) *Cart
	RemoveCart(itemidx int) bool
	CartClose()
}

func NewCartDBHandler(dbConn string) CartDBHandler {
	//handler = newCartMemHandler()
	return newCartpqHandler(dbConn)
}
