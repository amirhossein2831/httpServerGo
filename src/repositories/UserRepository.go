package repositories

import (
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"gorm.io/gorm"
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
	if res.Error != nil || res.RowsAffected == 0 {
		return user, errors.New("the id is not valid"), nil
	}
	return user, nil, res
}
