package Models

import (
	"errors"

	"alexedwards.net/CC-GO/Config"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type LoginCredentials struct {
	User     User
	Username string
	Password string
}

func CreateAUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func CheckLoginDetails(LoginCredentials LoginCredentials) (error, User) {
	var userModal User
	Config.DB.Where("email = ?", LoginCredentials.Username).First(&userModal)
	if userModal.Email == LoginCredentials.Username {
		error := bcrypt.CompareHashAndPassword([]byte(userModal.Password), []byte(LoginCredentials.Password))
		if error != nil {
			return error, userModal
		}

		return nil, userModal
	} else {
		error := errors.New("User name does not match")
		return error, userModal
	}
}
