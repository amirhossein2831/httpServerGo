package model

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	BirthData time.Time `json:"birth-data"`
	UserID    uint      `gorm:"foreignKey:UserID" json:"user-id"` // Foreign key

}

type User struct {
	gorm.Model
	FirstName string   `json:"first-name"`
	LastName  string   `json:"last-name"`
	Email     string   `json:"email"`
	Profile   *Profile `gorm:"foreignKey:UserID" json:"profile"`
}

//
//func initUser(firstName string, lastName string, email string, profile *Profile) *User {
//	return &User{FirstName: firstName, LastName: lastName, Email: email, Profile: profile}
//}
//
//func initProfile(age int, address string, birthDate time.Time) *Profile {
//	return &Profile{Age: age, Address: address, BirthData: birthDate}
//}
//
//func SeedUser() {
//	for i := 0; i < 10; i++ {
//		profile := initProfile(29, "iran", time.Time{})
//		user := initUser("amir", "motaghian", "amirmemool12@gmail.com", profile)
//
//		users = append(users, *user)
//	}
//}
//
//func GetUsers() []User {
//	return users
//}
//
//func GetUser(pk string) (User, error) {
//	id, err := strconv.Atoi(pk)
//	if err != nil {
//		return User{}, errors.New("invalid id")
//	}
//	for i := 0; i < len(users); i++ {
//		if users[i].ID == id {
//			return users[i], nil
//		}
//
//	}
//	return User{}, errors.New("cannot find user")
//}
//
//func CreateUser(user User) {
//	user.ID = len(users) + 1
//	users = append(users, user)
//}
//
//func UpdateUser(user User, body User) {
//	user.FirstName = body.FirstName
//	user.LastName = body.LastName
//	user.Email = body.Email
//	user.Profile = body.Profile
//
//	for i := 0; i < len(users); i++ {
//		if users[i].ID == user.ID {
//			users[i] = user
//		}
//	}
//}
//
//func RemoveUser(pk string) error {
//	id, err := strconv.Atoi(pk)
//	if err != nil {
//		return errors.New("invalid id")
//	}
//	for i := 0; i < len(users); i++ {
//		if users[i].ID == id {
//			users = append(users[:i], users[i+1:]...)
//			return nil
//		}
//	}
//	return errors.New("cannot find user")
//}
