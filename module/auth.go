package module

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Auth struct {
	ID       uint64 `gorm:"primary_key" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CheckAuth(email string, password string) (bool, error) {
	var account Account
	err := db.Select("id").Where("email = ? and password = ?", email, password).Find(&account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if account.ID > 0 {
		return true, nil
	}
	return false, nil
}

//检测用户是否存在
func ExistEmail(email string) (bool, error) {
	var account Account
	err := db.Select("id").Where("email = ?", email).Find(&account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if account.ID > 0 {
		return true, nil
	}

	return false, nil
}

//注册
func Register(email string, password string) (uint64, error) {

	var newId uint64
	db.Transaction(func(tx *gorm.DB) error {
		createAt := time.Now().Unix()
		account := Account{
			Email:    email,
			Password: password,
			CreateAt: createAt,
			Status:   1,
		}
		err := tx.Create(&account).Error
		if err != nil {
			return err
		}

		//member := AccountMember{
		//	Uid:      account.ID,
		//	CreateAt: createAt,
		//}
		//err = tx.Create(&member).Error
		//if err != nil {
		//	return err
		//}

		newId = account.ID
		return nil
	})

	return newId, nil

}
