package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
)

//注册
func Register(username string, phone string, email string, password string, ip string) (uint64, error) {

	var newId uint64

	db.Transaction(func(tx *gorm.DB) error {
		account := Account{
			Username:   username,
			Phone:      phone,
			Email:      email,
			Password:   password,
			CreateAt:   time.Now().Unix(),
			CreateIpAt: ip,
			Status:     1,
		}
		err := tx.Create(&account).Error
		if err != nil {
			return err
		}

		member := AccountMember{
			Uid:      account.ID,
			CreateAt: time.Now().Unix(),
		}
		err = tx.Create(&member).Error
		if err != nil {
			return err
		}

		newId = account.ID
		return nil
	})

	return newId, nil

}

//登录
func Login(title string, password string, ip string) error {
	db.Model(&Account{}).Updates(Account{
		LastLoginAt:   time.Now().Unix(),
		LastLoginIpAt: ip,
	}).Where()
}

//第三方登录
