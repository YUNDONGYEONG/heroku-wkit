package model

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/lib/pq" //암시적
)

type pqHandler struct {
	db *sql.DB // 멤버변수로 가진다
}

// Query문 가져오기
func (s *pqHandler) GetProducts() []*Product {
	products := []*Product{}                          //list를 만든다
	rows, err := s.db.Query("SELECT * FROM products") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var product Product                                                                                                                                                                                              //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&product.Idx, &product.SessionCode, &product.ContentName, &product.ProductName, &product.ContentMain, &product.File, &product.Price, &product.Registered, &product.CreateTime, &product.ProductStatus) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		products = append(products, &product)
	}
	//log.Print(products[0])
	return products
}

// Query문 추가
func (s *pqHandler) AddProduct(sessioncode string, contentname string, productname string, contentmain string, file string, price int, registered string) *Product { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db.Prepare("INSERT INTO products (sessioncode, contentname, productname, contentmain, file, price, registered, productstatus, createtime) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW()) RETURNING idx") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	var idx int
	err = stmt.QueryRow(sessioncode, contentname, productname, contentmain, file, price, registered, true).Scan(&idx)
	if err != nil {
		panic(err)
	}
	var product Product
	product.Idx = idx
	product.SessionCode = sessioncode
	product.ContentName = contentname
	product.ProductName = productname
	product.ContentMain = contentmain
	product.File = file
	product.Price = price
	product.Registered = registered
	product.CreateTime = time.Now()
	product.ProductStatus = true
	return &product
}

func (s *pqHandler) RemoveProduct(idx int) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := s.db.Prepare("DELETE FROM products WHERE idx=$1")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(idx)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

func (s *pqHandler) Detail(idx int) *Product {
	rows, err := s.db.Query("SELECT * FROM products WHERE idx = " + strconv.Itoa(idx)) //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var product Product
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.                                                                                                                                                  //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&product.Idx, &product.SessionCode, &product.ContentName, &product.ProductName, &product.ContentMain, &product.File, &product.Price, &product.Registered, &product.CreateTime, &product.ProductStatus) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
	}
	//log.Print(product)
	return &product
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (s *pqHandler) Close() {
	s.db.Close()
}

func newpqHandler(dbConn string) DBHandler {
	database, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		//file 컬럼 추가 하였습니다.
		`CREATE TABLE IF NOT EXISTS products (
			idx				SERIAL PRIMARY KEY,
			sessioncode		VARCHAR NOT NULL,
			contentname		VARCHAR(30) NOT NULL,
			productname		VARCHAR(30) NOT NULL,
			contentmain		VARCHAR(800) NOT NULL,
			file			bytea NOT NULL,
			price			INT NOT NULL,
			registered		VARCHAR(30),
			createtime		TIMESTAMP NOT NULL,
			productstatus 	BOOLEAN
			)`)
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	//임시로 데이터 넣기
	// statement, _ = database.Prepare(`INSERT INTO products (contentname, productname, contentmain, price, registered, productstatus,createtime) VALUES ('게시글제목','상품명','상품상세내용',1000,'등록자',true, datetime('now'))`)

	// statement.Exec()
	return &pqHandler{db: database} // &pqHandler{}를 반환
}
