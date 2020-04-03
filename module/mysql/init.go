package mysql

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init() {

	fmt.Println("------------- init sql")
	var err error
	db, err = gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/um_sys?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		os.Exit(-1)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&AccountPlatform{})
	db.AutoMigrate(&AccountMember{})
	db.AutoMigrate(&AccountStaff{})

}
