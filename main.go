package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// we make list of map there
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		//User inputes
		firstName, lastName, email, userTicket := getUserInput()

		// valitation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookingTickets(userTicket, firstName, lastName, email)

			// call func print Firsnames
			getFirstNames()
			if remainingTickets == 0 {
				fmt.Println("Our conferance if booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Firs name or last name you entered is to short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}
}

func greetUsers() {

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email addres:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookingTickets(userTicket uint, firstName, lastName, email string) {

	remainingTickets = remainingTickets - userTicket

	//create a map for user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["number of tickets"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %s\n", bookings)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTicket, email)

	fmt.Printf("%d tickets reamining for %s\n", remainingTickets, conferenceName)
}
