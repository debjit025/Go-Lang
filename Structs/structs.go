package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	// var appUser user
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		panic(err)
	}

	app := user.User{
		FirstName: "Max",
		LastName:  "Shwartz",
	}

	fmt.Println(app)

	// ... do something awesome with that gathered data!
	// outputUserDetails(*appUser)
	// appUser.ClearUserName()
	outputUserDetails(*appUser)
}

func outputUserDetails(u user.User) {
	fmt.Println(u)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
