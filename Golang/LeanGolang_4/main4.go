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

// danh sach data
type List struct {
	name               string `json:"name"`
	nickName           string `json:"nick_name"`
	age                int    `json:"age"`
	playingPosition    string `json:"playing_position"`
	teamCode           string `json:"team_code"`
	currentRank        string `json:"current_rank"`
	homeTown           string `json:"home_town"`
	identityCardNumber string `json:"identity_card_number"`
	phoneNumber        string `json:"phone_number"`
	accountName        string `json:"account_name"`
}

// inser len data
func InserData(db *sql.DB) {
	_, err := db.Exec("insert into danhsachdangky(ten, BietDanh, Tuoi, ViTriThiDau,MaDoi,RankHienTai,QueQuan,SoCmnd,Sdt,TenTaiKhoanGame) values(?,?,?,?,?,?,?,?,?,?)",
		"tuananh", "a", 25, "mid", "c1", "kimhcuong", "hn", 1234456678, 0544533423, "admin33")
	if err != nil {
		panic(err)
	}
	println("inser thanh cong")
}

// update data
func UpdateData(db *sql.DB) {
	_, err := db.Exec("update danhsachdangky set tuoi = 18 where tuoi = 25")
	if err != nil {
		panic(err)
	}
	println("Update thanh cong")
}

// select data
func Selectter(db *sql.DB) {
	stm, err := db.Query("select * from danhsachdangky")
	if err != nil {
		panic(err)
	}
	defer stm.Close()
	danhsach := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			panic(err)
		}
		danhsach = append(danhsach, d)
	}
	println(fmt.Sprintf("danh sach: %+v", danhsach))
}

func main() {
	var db *sql.DB
	// tao ket noi db
	db, err := sql.Open("mysql", "root:Admin@123@(localhost:3306)/tuananh")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	InserData(db)
	//UpdateData(db)
	//Selectter(db)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/v1/register", func(c echo.Context) error {
		Name := c.QueryParam("ten")
		nichName := c.QueryParam("bietDanh")
		age, err := strconv.ParseInt(c.QueryParam("tuoi"), 10, 64)
		if err != nil {
			return c.String(http.StatusOK, "loi1")
		}
		playingPosition := c.QueryParam("viTriThiDau")
		teamCode := c.QueryParam("maDoi")
		currentRank := c.QueryParam("rank")
		homeTown := c.QueryParam("queQuan")
		identityCardNumber := c.QueryParam("soCmnd")
		phoneNumber := c.QueryParam("soDienThoai")
		AccountName := c.QueryParam("taiKhoanGame")
		err = HienThiTatCaNguoiChoi(db, Name, nichName, int(age), playingPosition, teamCode, currentRank, homeTown, identityCardNumber, phoneNumber, AccountName)
		if err != nil {
			return c.String(http.StatusOK, "lỗi2")
		}
		return c.String(http.StatusOK, "thanh cong")
	})
	e.GET("/api/v1/get-by-age", func(c echo.Context) error {
		age := c.QueryParam("age")
		getAge, err := HienThiTheoTuoi(db, age)
		if err != nil {
			return c.String(http.StatusOK, "loi4")
		}
		return c.JSON(http.StatusOK, getAge)
	})
	e.GET("/api/v1/team", func(c echo.Context) error {
		teamCode := c.QueryParam("teamCode")
		getTeamCode, err := HienThiTheomaDoi(db, teamCode)
		if err != nil {
			return c.String(http.StatusOK, "loi")
		}
		return c.JSON(http.StatusOK, getTeamCode)
	})
	e.GET("/api/v1/rank", func(c echo.Context) error {
		rank := c.QueryParam("Rank")
		getRank, err := HienThiTheoRank(db, rank)
		if err != nil {
			return c.String(http.StatusOK, "loi")
		}
		return c.JSON(http.StatusOK, getRank)
	})
	e.GET("/api/v1/position", func(c echo.Context) error {
		playingPosition := c.QueryParam("playingPosition")
		getPlayingPosition, err := HienThiTheoViTri(db, playingPosition)
		if err != nil {
			return c.String(http.StatusOK, "loi")
		}
		return c.JSON(http.StatusOK, getPlayingPosition)
	})
	e.GET("/api/v1/phone-number", func(c echo.Context) error {
		phoneNumber := c.QueryParam("phoneNumber")
		getPhonenumber, err := TimThongTinQuaNumberPhone(db, phoneNumber)
		if err != nil {
			return c.String(http.StatusOK, "loi")
		}
		return c.JSON(http.StatusOK, getPhonenumber)
	})
	e.GET("/api/v1/identityCardNumber", func(c echo.Context) error {
		identityCardNumber := c.QueryParam("identityCardNumber")
		getIdentityCardNumber, err := TimThongTinQuaSoCmnd(db, identityCardNumber)
		if err != nil {
			return c.String(http.StatusOK, "loi")
		}
		return c.JSON(http.StatusOK, getIdentityCardNumber)
	})
	e.GET("/api/v1/username", func(c echo.Context) error {
		identityCardNumber := c.QueryParam("loi")
		getIdentityCardNumber, err := TimThongTinQuaTenTaiKhoan(db, identityCardNumber)
		if err != nil {
			return c.String(http.StatusOK, "loi")
		}
		return c.JSON(http.StatusOK, getIdentityCardNumber)
	})

	e.Start(":8888")
}

func HienThiTatCaNguoiChoi(db *sql.DB, name string, nichName string, age int, playingPosition string, teamCode string, currentRank string, homeTown string, identityCardNumber string, phoneNumber string, AccountName string) error {
	// sai rồi, insert cái thông tin của mình chứ k đc fix nhưu này
	//_, err := db.Exec("insert into danhsachdangky(Ten, BietDanh, Tuoi, ViTriThiDau,MaDoi,RankHienTai,QueQuan,SoCmnd,Sdt,TenTaiKhoanGame) values(?,?,?,?,?,?,?,?,?,?)",
	//	"nguyen", "lutkingofpain&tuoi", 25, "mid", "001", "thachdau", "namdinh", 1234456678, 0544533423, "kbop")

	_, err := db.Exec("insert into danhsachdangky(Ten, BietDanh, Tuoi, ViTriThiDau,MaDoi,RankHienTai,QueQuan,SoCmnd,Sdt,TenTaiKhoanGame) values(?,?,?,?,?,?,?,?,?,?)",
		name, nichName, age, playingPosition, teamCode, currentRank, homeTown, identityCardNumber, phoneNumber, AccountName)

	return err
}

func HienThiTheoTuoi(db *sql.DB, age string) ([]List, error) {
	stm, err := db.Query("select * from danhsachdangky where tuoi=?", age)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	list := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			return nil, err
		}
		list = append(list, d)
	}

	return list, nil
}

func HienThiTheomaDoi(db *sql.DB, teamCode string) ([]List, error) {
	stm, err := db.Query("select * from danhsachdangky where MADoi=?", teamCode)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	list := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			return nil, err
		}
		list = append(list, d)
	}

	return list, nil
}

func HienThiTheoRank(db *sql.DB, Rank string) ([]List, error) {
	stm, err := db.Query("select * from danhsachdangky where RankHienTai=?", Rank)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	list := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			return nil, err
		}
		list = append(list, d)
	}

	return list, nil
}

func HienThiTheoViTri(db *sql.DB, playingPosition string) ([]List, error) {
	stm, err := db.Query("select * from danhsachdangky where ViTriThiDau=?", playingPosition)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	list := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			return nil, err
		}
		list = append(list, d)
	}

	return list, nil
}

func TimThongTinQuaNumberPhone(db *sql.DB, phoneNumber string) ([]List, error) {
	stm, err := db.Query("select * from danhsachdangky where Sdt=?", phoneNumber)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	list := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			return nil, err
		}
		list = append(list, d)
	}

	return list, nil
}

func TimThongTinQuaSoCmnd(db *sql.DB, identityCardNumber string) ([]List, error) {
	// sai ten
	stm, err := db.Query("select * from danhsachdangky where SoCmnd=?", identityCardNumber)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	list := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			return nil, err
		}
		list = append(list, d)
	}

	return list, nil
}

func TimThongTinQuaTenTaiKhoan(db *sql.DB, AccountName string) ([]List, error) {
	stm, err := db.Query("select * from danhsachdangky where TenTaiKhoanGame=?", AccountName)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	list := make([]List, 0)
	for stm.Next() {
		d := List{}
		if err = stm.Scan(&d.name, &d.nickName, &d.age, &d.playingPosition, &d.teamCode, &d.currentRank, &d.homeTown, &d.identityCardNumber, &d.phoneNumber, &d.accountName); err != nil {
			return nil, err
		}
		list = append(list, d)
	}

	return list, nil
}
