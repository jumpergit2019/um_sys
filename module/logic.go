package module

import (
	"time"
)

//一般登录
func Login(title string, password string, ip string) error {
	rst := db.Model(&Account{}).Updates(Account{
		LastLoginAt:   time.Now().Unix(),
		LastLoginIpAt: ip,
	}).Where("password = ?", password).
		Where("email = ?", title).
		Or("phone = ?", title).
		Or("username = ?", title)

	if rst.Error != nil {
		return rst.Error
	}
	return nil
}

//todo:第三方登录
func PlatformLogin() {

}
