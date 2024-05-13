package repositories

import (
	"encoding/json"
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetUsers() ([]model.User, error, *gorm.DB) {
	var users []model.User
	res := config.App.GetDB().Find(&users)
	if res.Error != nil {
		return nil, res.Error, nil
	}
	return users, nil, res
}

func GetUser(pk string) (model.User, error, *gorm.DB) {
	var user model.User

	id, err := strconv.Atoi(pk)
	if err != nil {
		return user, errors.New("invalid id"), nil
	}

	res := config.App.GetDB().First(&user, id)
	if res.Error != nil {
		return user, res.Error, nil
	}
	return user, nil, res
}

func CreateUser(r *http.Request) (model.User, error, *gorm.DB) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return user, err, nil
	}
	res := config.App.GetDB().Create(&user)
	if res.Error != nil {
		return user, res.Error, nil
	}

	return user, nil, res
}

func DeleteUser(pk string) error {
	id, err := strconv.Atoi(pk)
	if err != nil {
		return err
	}
	res := config.App.GetDB().Delete(&model.User{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
