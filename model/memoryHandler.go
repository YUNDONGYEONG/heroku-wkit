package model

import (
	"time"
)

type memoryHandler struct {
	productMap map[int]*Product
}

//4개 func을 만든다
func (m *memoryHandler) GetProducts() []*Product {
	list := []*Product{}
	for _, v := range m.productMap {
		list = append(list, v)
	}
	return list
}

//
func (m *memoryHandler) AddProduct(sessioncode string, contentname string, productname string, contentmain string, file string, price int, registered string) *Product {
	idx := len(m.productMap) + 1
	createtime := time.Now()
	product := &Product{idx, sessioncode, contentname, productname, contentmain, file, price, registered, createtime, true}
	m.productMap[idx] = product
	return product
}

func (m *memoryHandler) RemoveProduct(id int) bool {
	if _, ok := m.productMap[id]; ok { // productMap id 값이 있으면
		delete(m.productMap, id) //지우고
		return true
	}
	return false
}

func (m *memoryHandler) Detail(idx int) *Product {
	product, _ := m.productMap[idx] // productMap id 값이 있으면
	return product
}

func (m *memoryHandler) Close() {

}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.productMap = make(map[int]*Product) // map을 초기화
	return m
}
