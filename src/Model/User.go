package Model

import (
	"time"
)

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

func CreateUser(id int, firstName string, lastName string, email string, profile *Profile) *User {
	return &User{ID: id, FirstName: firstName, LastName: lastName, Email: email, Profile: profile}
}

func createProfile(age int, address string, birthDate time.Time) *Profile {
	return &Profile{Age: age, Address: address, BirthData: birthDate}
}
