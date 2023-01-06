package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

// Student các thuộc tính phải có tên json trùng với tên trong bảng của db
type Student struct {
	Name      string `json:"name"`
	ClassName string `json:"className"`
	Age       int    `json:"age"`
	RollNo    string `json:"rollNo"`
}

// InsertExample Ví dụ câu lệnh insert
// thêm mới 1 dòng vào bảng
func InsertExample(db *sql.DB) {
	_, err := db.Exec("insert into student(name, className, age, rollNo) values(?,?,?,?)", "Nguyên", "Lớp 1", 25, "001")
	if err != nil {
		panic(err)
	}
	println("insert thanh cong!")
}

// SelectExample ví dụ câu lệnh select
// lấy các dòng trong bảng ra và in ra màn hình
func SelectExample(db *sql.DB) {
	rows, err := db.Query("select * from student")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	students := make([]Student, 0)
	for rows.Next() {
		s := Student{}
		if err = rows.Scan(&s.Name, &s.ClassName, &s.Age, &s.RollNo); err != nil {
			panic(err)
		}
		students = append(students, s)
	}

	println(fmt.Sprintf("students: %+v", students))
}

// UpdateExample ví dụ câu lệnh update
// cập nhật lại giá trị cột của 1 dòng nào đấy
func UpdateExample(db *sql.DB) {
	// cập nhật lại tuổi của những người có tuổi = 25 thành 100
	_, err := db.Exec("update student set age = 100 where age = 25")
	if err != nil {
		panic(err)
	}
	println("update thanh cong!")
}

func DeleteExample(db *sql.DB) {
	_, err := db.Exec("delete from student where age=21")
	if err != nil {
		panic(err)
	}
	println("delete thanh cong")
}

func InsertNewStudent(db *sql.DB, name string, className string, age int, rollNo string) error {
	_, err := db.Exec("insert into student(name, className, age, rollNo) values(?,?,?,?)",
		name, className, age, rollNo)
	return err
}

func GetStudentByRollNo(db *sql.DB, rollNo string) ([]Student, error) {
	rows, err := db.Query("select * from student where rollNo = ?", rollNo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	students := make([]Student, 0)
	for rows.Next() {
		s := Student{}
		if err = rows.Scan(&s.Name, &s.ClassName, &s.Age, &s.RollNo); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

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
	//InsertExample(db)
	//SelectExample(db)
	////DeleteEcample(db)
	//UpdateExample(db)
	//SelectExample(db)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/create-student", func(c echo.Context) error {
		name := c.QueryParam("name")
		className := c.QueryParam("className")
		age, err := strconv.ParseInt(c.QueryParam("age"), 10, 64)
		if err != nil {
			return c.String(http.StatusOK, "lỗi")
		}
		rollNo := c.QueryParam("rollNo")
		err = InsertNewStudent(db, name, className, int(age), rollNo)
		if err != nil {
			return c.String(http.StatusOK, "lỗi")
		}
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/get-student", func(c echo.Context) error {
		rollNo := c.QueryParam("rollNo")
		students, err := GetStudentByRollNo(db, rollNo)
		if err != nil {
			return c.String(http.StatusOK, "lỗi")
		}
		return c.JSON(http.StatusOK, students)
	})

	e.Start(":1234")
}
