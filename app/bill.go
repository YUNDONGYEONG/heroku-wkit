package app

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

// 2021-11-05, 영수증 인쇄시 키 값으로 사용할 uuid 생성, 전송, 장재혁
func (a *AppHandler) getBillKeyHandler(w http.ResponseWriter, r *http.Request) {
	bill_key := uuid.New()
	log.Println(bill_key.String())

	rd.JSON(w, http.StatusOK, bill_key)
}
