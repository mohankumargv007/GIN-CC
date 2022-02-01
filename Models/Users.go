package Models

import (
	"alexedwards.net/CC-GO/Config"
	_ "github.com/go-sql-driver/mysql"
)

func CreateAUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
