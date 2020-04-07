package module

import (
	"fmt"
	"os"
	"um_sys/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup() {

	fmt.Println("------------- init sql")
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
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
