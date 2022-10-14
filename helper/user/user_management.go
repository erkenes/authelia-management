package user

import (
	"authelia-users/helper/basics"
	"fmt"
	"github.com/imdario/mergo"
	"sort"
)

// CreateUser Create a new user
func CreateUser(username string, password string, email string, groups []string, displayName string) {
	currentUsers := getCurrentUsers()
	newUser := createNewUser(username, password, email, groups, displayName)

	if checkIfUserExists(username) {
		fmt.Println(basics.ColorGreen + "Updated the user `" + username + "` successfully." + basics.ColorReset)
	} else {
		fmt.Println(basics.ColorGreen + "Created new user `" + username + "` successfully." + basics.ColorReset)
	}

	mergo.Merge(&newUser, currentUsers)

	saveToUserFile(newUser)
}

// RemoveUser Remove a specific user
func RemoveUser(username string) {
	currentUsers := getCurrentUsers()

	if checkIfUserExists(username) {
		delete(currentUsers["users"], username)

		saveToUserFile(currentUsers)
		fmt.Println(basics.ColorGreen + "Removed the user `" + username + "` successfully." + basics.ColorReset)
	} else {
		fmt.Println(basics.ColorRed + "User `" + username + "` not found!" + basics.ColorReset)
	}
}

// AddGroupToUser Add a group to a user
func AddGroupToUser(username string, groups []string) {
	currentUsers := getCurrentUsers()

	if checkIfUserExists(username) {
		userData := currentUsers["users"][username]

		groups = basics.UniqueStrings(currentUsers["users"][username].Groups, groups)
		sort.Strings(groups)

		newUserData := UserData{
			Groups: groups,
		}

		mergo.Merge(&newUserData, userData)

		currentUsers["users"][username] = newUserData

		saveToUserFile(currentUsers)

		fmt.Println(basics.ColorGreen + "Updated the user `" + username + "` successfully." + basics.ColorReset)
	} else {
		fmt.Println(basics.ColorRed + "User `" + username + "` not found!" + basics.ColorReset)
	}
}

// RemoveGroupFromUser Remove a group from a user
func RemoveGroupFromUser(username string, group string) {
	currentUsers := getCurrentUsers()

	if checkIfUserExists(username) {
		if checkIfUserHasGroup(username, group) {
			userData := currentUsers["users"][username]

			userGroups := userData.Groups
			userGroups = basics.RemoveFromStringSlice(userGroups, group)

			newUserData := UserData{
				Groups: userGroups,
			}

			mergo.Merge(&newUserData, userData)

			currentUsers["users"][username] = newUserData

			saveToUserFile(currentUsers)

			fmt.Println(basics.ColorGreen + "Removed the group `" + group + "` from the `" + username + "` successfully." + basics.ColorReset)
		} else {
			fmt.Println(basics.ColorYellow + "The user `" + username + "` is not a member of the group `" + group + "`!" + basics.ColorReset)
		}
	} else {
		fmt.Println(basics.ColorRed + "User `" + username + "` not found!" + basics.ColorReset)
	}
}

// CheckIfUserExists Check if a user exists
func CheckIfUserExists(username string) bool {
	return checkIfUserExists(username)
}
