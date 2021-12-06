package app

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *AppHandler) getCartListHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSesssionID(r)
	//log.Print("getcart", sessionId)

	list := a.cartdb.GetCarts(sessionId)
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addCartHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSesssionID(r)
	idx, _ := strconv.Atoi(r.FormValue("idx"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	cart := a.cartdb.AddCart(idx, quantity, sessionId)
	rd.JSON(w, http.StatusCreated, cart)
}

func (a *AppHandler) removeCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idx, _ := strconv.Atoi(vars["idx"])
	ok := a.cartdb.RemoveCart(idx)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) CartClose() {
	a.cartdb.CartClose()
}
