package helper

import "strings"

func ValidateUserInput(firstName, lastName, email string, userTicket, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}
