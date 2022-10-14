package basics

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"net/mail"
	"strconv"
	"syscall"
)

// GetTextInput get a text input
func GetTextInput(question string, required bool) string {
	var output string

	if question != "" {
		fmt.Println(question)
	}

	fmt.Printf("> " + ColorGreen)
	fmt.Scanln(&output)
	fmt.Println(ColorReset)

	if required == true && output == "" {
		fmt.Println(ColorRed + "An empty value is not allowed here" + ColorReset)
		output = GetTextInput(question, required)
	}

	return output
}

// GetNumberInput get a number input
func GetNumberInput(question string, required bool) int {
	var output int
	var err error

	text := GetTextInput(question, required)
	output, err = strconv.Atoi(text)
	if err != nil {
		fmt.Println(ColorRed + "Please enter a valid number" + ColorReset)
		output = GetNumberInput(question, required)
	}

	return output
}

// GetConfirmInput ask for a confirmation
func GetConfirmInput(question string) bool {
	var output bool

	if question != "" {
		fmt.Println(question)
	}

	output = askForConfirmation(true)
	fmt.Println(ColorReset)

	return output
}

// GetEmailAddressInput get an email address input and validate it
func GetEmailAddressInput(question string, required bool) string {
	var output string

	if question != "" {
		fmt.Println(question)
	}

	fmt.Printf("> " + ColorGreen)
	fmt.Scanln(&output)
	fmt.Println(ColorReset)

	if required == true && output == "" {
		fmt.Println(ColorRed + "An empty value is not allowed here" + ColorReset)
		output = GetEmailAddressInput(question, required)
	}

	_, err := mail.ParseAddress(output)

	if err != nil {
		fmt.Println(ColorRed + "The email address is not valid. Please try again." + ColorReset)
		output = GetEmailAddressInput(question, required)
	}

	return output
}

// GetPasswordInput get a hidden password input
func GetPasswordInput(question string, required bool) string {
	var output string
	var passwd []byte
	var err error

	if question != "" {
		fmt.Println(question)
	}

	fmt.Printf("> " + ColorGreen)
	passwd, err = terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println(ColorReset)

	if err != nil {
		fmt.Println(ColorRed + "An error occurred. Please try again." + ColorReset)
		output = GetPasswordInput(question, required)
	} else {
		output = string(passwd)
	}

	if required == true && output == "" {
		fmt.Println(ColorRed + "An empty value is not allowed here" + ColorReset)
		output = GetPasswordInput(question, required)
	}

	return output
}

// PrintSectionHeader Print a section header
func PrintSectionHeader(title string) {
	fmt.Println("\n\n\n --- " + ColorCyan + title + ColorReset + " --- \n")
}

// AskForConfirmation ask for a confirmation
func askForConfirmation(withPrefix bool) bool {
	var response string
	if withPrefix == true {
		fmt.Printf("> " + ColorGreen)
	}
	fmt.Scanln(&response)
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println(ColorRed + "Please type yes or no and then press enter:" + ColorReset)
		return askForConfirmation(withPrefix)
	}
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}
