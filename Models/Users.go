package Models

import (
	"errors"

	"alexedwards.net/CC-GO/Config"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateAUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func CheckLoginDetails(user *User, username string, password string) (err error, Id int, Name string, Role string) {
	var userModal User
	Config.DB.Where("email = ?", username).First(&userModal)
	if userModal.Email == username {
		err = bcrypt.CompareHashAndPassword([]byte(userModal.Password), []byte(password))
		if err != nil {
			return err, 0, "", ""
		}
		return nil, int(userModal.ID), userModal.FirstName + "" + userModal.LastName, userModal.Role
	} else {
		err := errors.New("User name does not match")
		return err, 0, "", ""
	}
}
