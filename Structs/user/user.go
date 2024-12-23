package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	birthdate string
	createdAt time.Time
}

func (u User) OutputUserDetails() { //Attaching a function to a struct
	fmt.Println(u.FirstName, u.LastName, u.createdAt)
}

func (u *User) ClearUserName() { //We have to use the pointer here to change the original values, otherwise we are just changing the copies.
	u.FirstName = ""
	u.LastName = ""
}

func NewUser(userFirstName, userLastName, userBirthdate string) (*User, error) { //can also use *user
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("Firstname, lastname and birthdate are required")
	}
	return &User{ //Can also use &user to save memory
		FirstName: userFirstName, //Can remove the first words if the sequence is the same
		LastName:  userLastName,  //Can remove the first words if the sequence is the same
		birthdate: userBirthdate, //Can remove the first words if the sequence is the same
		createdAt: time.Now(),    //Can remove the first words if the sequence is the same
	}, nil
}
