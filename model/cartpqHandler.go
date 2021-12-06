package model

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type cartpqHandler struct {
	db *sql.DB
}

//GetCart 인자로 sessionId를 받고
func (s *cartpqHandler) GetCarts(sessionId string) []*CartPrdt {
	cartprdts := []*CartPrdt{}
	//WHERE 절로 sessionId=?, sessionId 인자추가하여 아래 DB table에 Index를 걸어주는 부분으로 매우 중요
	rows, err := s.db.Query("SELECT idx, sessioncode, quantity, addtime, contentname, registered, file, price FROM products INNER JOIN carts ON carts.productidx = products.idx WHERE sessionId=$1", sessionId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cartprdt CartPrdt
		rows.Scan(&cartprdt.Product.Idx, &cartprdt.Product.SessionCode, &cartprdt.Quantity, &cartprdt.Addtime, &cartprdt.Product.ContentName, &cartprdt.Product.Registered, &cartprdt.Product.File, &cartprdt.Product.Price)
		// cartprdt.Cart.Idx = cartprdt.Product.Idx
		cartprdts = append(cartprdts, &cartprdt)
	}
	//log.Print(cartprdts[0])
	return cartprdts
}

//AddCart func에서 name뿐만 아니라, sessionId까지 받아준다.
func (s *cartpqHandler) AddCart(productidx int, quantity int, sessionId string) *Cart { //?, ?, ?인자가 세 개가 된다
	stmt, err := s.db.Prepare("INSERT INTO carts (productidx, quantity, sessionId, addtime) VALUES ($1, $2, $3, NOW()) ON CONFLICT (productidx, sessionId) DO UPDATE SET quantity = $2")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(productidx, quantity, sessionId)
	if err != nil {
		panic(err)
	}
	var cart Cart
	cart.ProductIdx = productidx
	cart.Quantity = quantity
	cart.Addtime = time.Now()
	return &cart
}

func (s *cartpqHandler) RemoveCart(productidx int) bool {
	stmt, err := s.db.Prepare("DELETE FROM carts WHERE productidx=$1")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(productidx)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

func (s *cartpqHandler) CartClose() {
	s.db.Close()
}

func newCartpqHandler(dbConn string) CartDBHandler {
	database, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}
	//sessionId 컬럼을 추가 지정해주어야 세션Id 별로 다르게 접속할 수 있다.
	//기본키를 (productidx, sessionId) 두 컬럼으로 설정 같은 세션아이디 일 경우 상품 중복 방지
	//productidx를 외래키로 설정해보았으나 작동 여부 불분명
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS carts (
			productidx  INT NOT NULL,
			quantity 	INT NOT NULL,
			sessionId 	VARCHAR(256) NOT NULL,
			addtime 	TIMESTAMP NOT NULL,
			PRIMARY KEY(productidx, sessionId),
			CONSTRAINT productidx_fk FOREIGN KEY(productidx) REFERENCES products(idx) on delete cascade
		);`)
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}
	statement, err = database.Prepare(
		`CREATE INDEX IF NOT EXISTS sessionIdIndexOnTodos ON carts (
				sessionId ASC
			);`)
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}
	return &cartpqHandler{db: database}
}

//CREATE INDEX~ EXISTS 인덱스 이름(sessionIdIndexOnCarts)을 해주면 서칭키를 제공해서 WHERE절로 더 빨리 찾는다
//새로운 컬럼의 인덱스를 할 때는 새로운 인덱스를 만들어 준다
//새로운 인덱스를 만들 때는 CREATE 구문 끝에 ~~); 괄호 뒤 세미콜론을 해주어서 statement 구분을 한다.
//sessionIdIndexOnCarts는 Index이름이다
//~~ON(어디에 있는 인덱스인가?) carts(테이블이름)에 있는 인덱스다
//소괄호 이후 어떤 컬럼에 인덱스를 걸 것 인가? sessionId
//어떤 방식으로 서칭트리를 만들 건가? ASC는 순차적 정렬을 하겠다는 의미
