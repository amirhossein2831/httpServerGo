package repositories

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func GetUsers() ([]model.User, error) {
	var users []model.User
	res := config.App.GetDB().Preload("Profile").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func GetUser(id string) (model.User, error) {
	var user model.User

	res := config.App.GetDB().First(&user, id)
	if res.Error != nil {
		return model.User{}, res.Error
	}
	return user, nil
}

func CreateUser(r *http.Request) (model.User, error) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hashedPassword)
	res := config.App.GetDB().Create(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}

func UpdateUser(r *http.Request, id string) (model.User, error) {
	var user model.User
	err := config.App.GetDB().First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	config.App.GetDB().Save(&user)

	return user, nil
}

func DeleteUser(id string) error {
	if err := config.App.GetDB().Where("user_id = ?", id).Delete(&model.Profile{}).Error; err != nil {
		return err
	}

	if err := config.App.GetDB().Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
