package user

import (
	"authelia-users/helper/basics"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
)

type UserData struct {
	DisplayName string   `yaml:"displayname"`
	Password    string   `yaml:"password"`
	Email       string   `yaml:"email"`
	Groups      []string `yaml:"groups"`
}

// getCurrentUsers Get all users
func getCurrentUsers() map[string]map[string]UserData {
	content := basics.ReadFile(".", "users_database.yml")

	var err error
	m := make(map[string]map[string]UserData)

	err = yaml.Unmarshal(content, &m)
	if err != nil {
		panic(err)
	}

	return m
}

// createNewUser Create a new user
func createNewUser(username string, password string, email string, groups []string, displayName string) map[string]map[string]UserData {
	// Establish the parameters to use for Argon2.
	p := &Argon2Params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	hash, err := generateFromPassword(password, p)
	if err != nil {
		log.Fatal(err)
	}

	newUser := map[string]map[string]UserData{
		"users": {
			username: {
				DisplayName: displayName,
				Password:    hash,
				Email:       email,
				Groups:      groups,
			},
		},
	}

	return newUser
}

// saveToUserFile Save the user file
func saveToUserFile(users map[string]map[string]UserData) {
	yamlData, err := yaml.Marshal(users)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	basics.WriteFile(".", "users_database.yml", yamlData, 0644)
}

// checkIfUserExists Check if a user exists
func checkIfUserExists(username string) bool {
	currentUsers := getCurrentUsers()

	if _, ok := currentUsers["users"][username]; ok {
		return true
	}

	return false
}

// checkIfUserHasGroup Check if a user is a member of a specific group
func checkIfUserHasGroup(username string, group string) bool {
	currentUsers := getCurrentUsers()

	currentGroups := currentUsers["users"][username].Groups

	return basics.StringSliceContains(currentGroups, group)
}
