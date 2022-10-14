package main

import (
	"authelia-users/helper/basics"
	"authelia-users/helper/user"
	"fmt"
)

var appVersion = "1.0.0"

func main() {
	fmt.Println("#################")
	fmt.Println("")
	fmt.Println("App version: " + basics.ColorCyan + appVersion + basics.ColorReset)
	fmt.Println("")
	fmt.Println("#################")
	home()
}

// home the home screen
func home() {
	basics.PrintSectionHeader("What do you want to do?")
	basics.PrintSectionHeader("User-Management")
	fmt.Println("1: Create a new user")
	fmt.Println("2: Delete a user")
	fmt.Println("3: Add a group to a user")
	fmt.Println("4: Remove a group from a user")

	inputMenu := basics.GetNumberInput("", true)

	switch inputMenu {
	case 1:
		openCreateNewUser()
		break
	case 2:
		openDeleteUser()
		break
	case 3:
		openAddGroupToUser()
		break
	case 4:
		openRemoveGroupFromUser()
		break
	}

	home()
}

// openCreateNewUser create a new user
func openCreateNewUser() {
	username := basics.GetTextInput("What is the username?", true)

	if user.CheckIfUserExists(username) {
		createNew := basics.GetConfirmInput(basics.ColorRed + "There is already a user with the name `" + username + "`. What should we do?" + basics.ColorReset)

		if !createNew {
			fmt.Printf(basics.ColorYellow + "Aborted." + basics.ColorReset)
			return
		}
	}

	password := basics.GetPasswordInput("Enter a password", true)
	email := basics.GetEmailAddressInput("Enter an email address", true)
	groups := askForGroups()
	displayName := basics.GetTextInput("What is the full name of the user?", true)

	user.CreateUser(username, password, email, groups, displayName)
}

// askForGroups Aks repeatedly which groups should be added
func askForGroups() []string {
	var groups []string
	group := basics.GetTextInput("Which group should be added? Leave empty to continue.", false)
	if group != "" {
		groups = append(groups, group)
		groups = append(groups, askForGroups()...)
	} else {
		fmt.Printf(basics.ColorYellow + "Skipped..." + basics.ColorReset)
	}

	return groups
}

// openDeleteUser delete a specific user
func openDeleteUser() {
	username := basics.GetTextInput("What is the username of the user that should be removed?", true)

	if user.CheckIfUserExists(username) {
		confirm := basics.GetConfirmInput("Do you really want to remove the user `" + username + "`? This can not be undone!")

		if confirm {
			user.RemoveUser(username)
		} else {
			fmt.Printf(basics.ColorYellow + "Aborted." + basics.ColorReset)
		}
	} else {
		fmt.Printf(basics.ColorRed + "The user `" + username + "` can not be found!" + basics.ColorReset)
	}
}

// openAddGroupToUser add a group to a user
func openAddGroupToUser() {
	username := basics.GetTextInput("What is the username of the user?", true)

	if user.CheckIfUserExists(username) {
		groups := askForGroups()

		user.AddGroupToUser(username, groups)
	} else {
		fmt.Printf(basics.ColorRed + "The user `" + username + "` can not be found!" + basics.ColorReset)
	}
}

// openRemoveGroupFromUser remove a group from a user
func openRemoveGroupFromUser() {
	username := basics.GetTextInput("What is the username of the user?", true)

	if user.CheckIfUserExists(username) {
		group := basics.GetTextInput("Which group should be removed?", true)

		user.RemoveGroupFromUser(username, group)
	} else {
		fmt.Printf(basics.ColorRed + "The user `" + username + "` can not be found!" + basics.ColorReset)
	}
}
