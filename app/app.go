package app

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tuckersWeb/wkit/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var rd *render.Render = render.New() //전역변수 render.New() 초기화

type AppHandler struct {
	http.Handler //handler http.Handler인데 handler를 생략, 암시적으로 인터페이스를 포함한 멤버 변수를 포함한 상태
	db           model.DBHandler
	cartdb       model.CartDBHandler
}

//hash 만드는 함수
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getSesssionID(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	// Set some session values.
	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list.html", http.StatusTemporaryRedirect)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) getProductListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetProducts() //model -> a.db로 바꾼다
	//log.Print(list)
	rd.JSON(w, http.StatusOK, list)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) addProductHandler(w http.ResponseWriter, r *http.Request) { //product list add 해주는 핸들러

	//Idx := r.FormValue("idx")

	Contentname := r.FormValue("contentname")
	Productname := r.FormValue("productname")
	Contentmain := r.FormValue("contentmain")
	uploadFile := r.FormValue("file")
	//log.Println("file", uploadFile)
	Price, _ := strconv.Atoi(r.FormValue("price"))
	Registered := r.FormValue("registered")
	SessionCode := GetMD5Hash(Productname)
	//Createtime := r.FormValue("mobile")
	// Productstatus := r.FormValue("")
	product := a.db.AddProduct(SessionCode, Contentname, Productname, Contentmain, uploadFile, Price, Registered)
	//log.Print(Contentname, Productname, Contentmain, Price, Registered)
	rd.JSON(w, http.StatusCreated, product) // JSON으로 product 값을 반환
}

type Success struct { //(클라이언트) 응답 결과를 알려주기 위한 구조체
	Success bool `json:"success"`
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) removeProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idx, _ := strconv.Atoi(vars["idx"]) // id값을 가져온다
	ok := a.db.RemoveProduct(idx)       //model -> a.db로 바꾼다
	if ok {
		rd.JSON(w, http.StatusOK, Success{true}) //ok 성공시
		// a.removeCartHandler(w, r)
	} else { //없는 경우
		rd.JSON(w, http.StatusOK, Success{false}) // 실패를 알려준다
	}
}

func (a *AppHandler) detailProductHandler(w http.ResponseWriter, r *http.Request) {
	//log.Print("dddddd")
	vars := mux.Vars(r)
	idx, _ := strconv.Atoi(vars["idx"])
	//log.Print(idx)
	detail := a.db.Detail(idx) //model -> a.db로 바꾼다
	//log.Print(detail)
	rd.JSON(w, http.StatusOK, detail)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) Close() { //새롭게 Close()를 외부에서 만들어 준 것.
	a.db.Close() //model -> a.db로 바꾼다
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// if request URL is /signin.html, then next()
	if strings.Contains(r.URL.Path, "/signin") ||
		strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	// if user already signed in
	sessionID := getSesssionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}

	// if not user sign in
	// redirect singin.html
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func MakeHandler(dbConn string) *AppHandler { //http.Handler를 *AppHandler로 변환시켜, NewHandler할 때 AppHandler를 초기화시켜 준다

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(CheckSignin),
		negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)
	a := &AppHandler{
		Handler: n, //r mux.Router()
		db:      model.NewDBHandler(dbConn),
		cartdb:  model.NewCartDBHandler(dbConn),
	}
	//일반 핸들러함수가 아니라 메소드로 바뀌었으니 a메소드로 넘겨준다
	//테스트용 핸들러
	//r.HandleFunc("/JUNlist", a.getProductListHandler).Methods("GET")
	r.HandleFunc("/list", a.getProductListHandler).Methods("GET")
	r.HandleFunc("/detail/{idx:[0-9]+}", a.detailProductHandler).Methods("GET")
	r.HandleFunc("/upload", a.addProductHandler).Methods("POST")
	r.HandleFunc("/detail/{idx:[0-9]+}", a.removeProductHandler).Methods("DELETE")
	r.HandleFunc("/cart", a.getCartListHandler).Methods("GET")
	r.HandleFunc("/add_cart", a.addCartHandler).Methods("POST")
	r.HandleFunc("/cart/{idx:[0-9]+}", a.removeCartHandler).Methods("DELETE")
	r.HandleFunc("/bill", a.getBillKeyHandler).Methods("GET") // 2021-11-05, /bill 경로 GET 방식일 경우 uuid 생성, 전송하는 getBillKeyHandler 실행, 장재혁
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)
	r.HandleFunc("/", a.indexHandler)

	return a // a 메소드를 반환
}
