package model

import (
	"time"
)

type Product struct {
	Idx           int       `json:"idx"`
	SessionCode   string    `json:"sessioncode"`
	ContentName   string    `json:"contentname"`
	ProductName   string    `json:"productname"`
	ContentMain   string    `json:"contentmain"`
	File          string    `json:"file"`
	Price         int       `json:"price"`
	Registered    string    `json:"registered"`
	CreateTime    time.Time `json:"createtime"`
	ProductStatus bool      `json:"productstatus"`
}

type DBHandler interface {
	GetProducts() []*Product
	AddProduct(sessioncode string, contentname string, productname string, contentmain string, file string, price int, registered string) *Product
	RemoveProduct(idx int) bool
	Detail(idx int) *Product
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

func NewDBHandler(dbConn string) DBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	//handler - newMemoryHandler()
	return newpqHandler(dbConn)
}
