package Model

import (
	"time"
)

var users []User

type Profile struct {
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	BirthData time.Time `json:"birth-data"`
}

type User struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first-name"`
	LastName  string   `json:"last-name"`
	Email     string   `json:"email"`
	Profile   *Profile `json:"profile"`
}

func initUser(id int, firstName string, lastName string, email string, profile *Profile) *User {
	return &User{ID: id, FirstName: firstName, LastName: lastName, Email: email, Profile: profile}
}

func initProfile(age int, address string, birthDate time.Time) *Profile {
	return &Profile{Age: age, Address: address, BirthData: birthDate}
}

func SeedUser() {
	for i := 0; i < 10; i++ {
		profile := initProfile(29, "iran", time.Time{})
		user := initUser(i+1, "amir", "motaghian", "amirmemool12@gmail.com", profile)

		users = append(users, *user)
	}
}

func GetUsers() []User {
	return users
}

