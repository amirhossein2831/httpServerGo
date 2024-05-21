package repositories

import (
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) All() ([]model.User, error) {
	var users []model.User

	err := DB.GetInstance().GetDb().Preload("Profile").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) Get(id string) (model.User, error) {
	var user model.User

	err := DB.GetInstance().GetDb().First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetByColumn(email string) (model.User, error) {
	var user model.User

	err := DB.GetInstance().GetDb().Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Create(user model.User) (model.User, error) {
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

func (ur *UserRepository) Update(user model.User, id string) (model.User, error) {
	err := ur.HardDelete(id)
	if err != nil {
		return model.User{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user.Password = string(hashedPassword)
	Id, _ := strconv.Atoi(id)
	user.ID = uint(Id)
	err = DB.GetInstance().GetDb().Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) SoftDelete(id string) error {
	if err := DB.GetInstance().GetDb().Where("user_id = ?", id).Delete(&model.Profile{}).Error; err != nil {
		return err
	}

	if err := DB.GetInstance().GetDb().Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) HardDelete(id string) error {
	if err := DB.GetInstance().GetDb().Unscoped().Where("user_id = ?", id).Delete(&model.Profile{}).Error; err != nil {
		return err
	}

	if err := DB.GetInstance().GetDb().Unscoped().Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
