package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var db *sql.DB
	// khai báo biến db là một con trỏ đến sql.DB

	// kết nối đến cơ sở dữ liệu MySQL
	db, err := sql.Open("mysql", "root:Admin@123@(localhost:3306)/tuananh")
	if err != nil {
		// xử lý lỗi nếu có
		panic(err.Error())
	}

	defer db.Close() // đóng kết nối sau khi sử dụng xong
	// Create
	_, err = db.Query("INSERT INTO Student(name,classNAme,age,rollNo) VALUES(?,?,?,?)", "ta", "a001", 21, "ydye")
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("kgjfdkglf"))
}
