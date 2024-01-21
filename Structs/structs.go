package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// func (u user) outputUserDetails() { //Attaching a function to a struct
// 	fmt.Println(u.firstName, u.lastName, u.createdAt)
// }

func (u *user) clearUserName() { //We have to use the pointer here to change the original values, otherwise we are just changing the copies.
	u.firstName = ""
	u.lastName = ""
}

func newUser(userFirstName, userLastName, userBirthdate string) user { //can also use *user
	return user{ //Can also use &user to save memory
		firstName: userFirstName, //Can remove the first words if the sequence is the same
		lastName:  userLastName,  //Can remove the first words if the sequence is the same
		birthdate: userBirthdate, //Can remove the first words if the sequence is the same
		createdAt: time.Now(),    //Can remove the first words if the sequence is the same
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user
	appUser := newUser(userFirstName, userLastName, userBirthdate)

	// ... do something awesome with that gathered data!
	outputUserDetails(appUser)
	appUser.clearUserName()
	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
