package user

import (
	"errors"
	"fmt"
	"time"
)

// lower case for this type as its not used outside of this package
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// we need the *user pointer here so the original user struct gets mutated and not a copy of it.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (Admin, error) {
	if email == "" || password == "" {
		return Admin{}, errors.New("Email and password are required")
	}

	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}, nil
}
