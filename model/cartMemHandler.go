package model

import "time"

type cartMemHandler struct {
	cartMap    map[int]*Cart
	cartmemMap map[int]*CartPrdt
}

//sessionId 인자추가
func (m *cartMemHandler) GetCarts(sessionId string) []*CartPrdt {
	list := []*CartPrdt{}
	for _, v := range m.cartmemMap {
		list = append(list, v)
	}
	return list
}

//sessionId 인자추가
func (m *cartMemHandler) AddCart(productidx int, quantity int, sessionId string) *Cart { //Table대로 sessionId가 먼저, 그 다음 name이 와야한다.
	cart := &Cart{productidx, quantity, time.Now()}
	m.cartMap[productidx] = cart
	return cart
}

func (m *cartMemHandler) RemoveCart(productidx int) bool {
	if _, ok := m.cartMap[productidx]; ok {
		delete(m.cartMap, productidx)
		return true
	}
	return false
}

func (m *cartMemHandler) CartClose() {

}

func newCartMemHandler() CartDBHandler {
	m := &cartMemHandler{}
	m.cartMap = make(map[int]*Cart)
	m.cartmemMap = make(map[int]*CartPrdt)
	return m
}
