package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

type name struct {
	NameTg      string `json:"nameTg"`
	Nametacpham string `json:"nametacpham"`
	Age         int    `json:"age"`
}

func Inser(dc *sql.DB) {
	_, err := dc.Exec("insert into books(nameTg, nametacpham, age) values(?,?,?)", "tuan anh", "ax", 22)
	if err != nil {
		log.Fatal(err)
	}
	println("Inser thanhf cong")
}
func Update(dc *sql.DB) {
	_, err := dc.Exec("update books set age = 25 where age = 22")
	if err != nil {
		log.Fatal(err)
	}
	println("update thanh cong")
}
func Selectter(dc *sql.DB) {
	row, err := dc.Query("select * from books")
	if err != nil {
		panic(err)
	}
	defer row.Close()
	books := make([]name, 0)
	for row.Next() {
		b := name{}
		if err = row.Scan(&b.NameTg, &b.Nametacpham, &b.Age); err != nil {
			log.Fatal(err)
		}
		books = append(books, b)
	}
	println(fmt.Sprintf("books: %+v", books))
}
func main() {
	var dc *sql.DB
	dc, err := sql.Open("mysql", "root:Admin@123@(localhost:3306)/ifdg")
	if err != nil {
		log.Fatal(err)
	}
	defer dc.Close()
	//e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//e.GET("/name", func(c echo.Context) error {
	//
	//})
	Inser(dc)
	Update(dc)
	Selectter(dc)
}
