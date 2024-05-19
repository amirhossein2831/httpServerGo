package repositories

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func GetUsers() ([]model.User, error) {
	var users []model.User

	err := DB.GetInstance().GetDb().Preload("Profile").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id string) (model.User, error) {
	var user model.User

	err := DB.GetInstance().GetDb().First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User

	err := DB.GetInstance().GetDb().Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
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
	err = DB.GetInstance().GetDb().Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func UpdateUser(r *http.Request, id string) (model.User, error) {
	var user model.User

	err := DB.GetInstance().GetDb().First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return model.User{}, err
	}

	DB.GetInstance().GetDb().Save(&user)

	return user, nil
}

func DeleteUser(id string) error {
	if err := DB.GetInstance().GetDb().Where("user_id = ?", id).Delete(&model.Profile{}).Error; err != nil {
		return err
	}

	if err := DB.GetInstance().GetDb().Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
