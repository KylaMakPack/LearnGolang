package helper

import "strings"

// need to capatilize to use package in another
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	//check first and last name 2 or more char
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// email has to contain @
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
