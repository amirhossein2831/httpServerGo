package Model

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Email     string
	BirthData time.Time
}

func CreateUser(id int, firstName string, lastName string, age int, email string, BirthData time.Time) *User {
	return &User{ID: id, FirstName: firstName, LastName: lastName, Age: age, Email: email, BirthData: BirthData}
}
