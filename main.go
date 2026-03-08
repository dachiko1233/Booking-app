package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// slice and Array
	var bookings []string

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	for {
		var firstName string
		var lastname string
		var email string
		var userTicket uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastname)

		fmt.Println("Enter your email addres:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTicket)
		if userTicket > remainingTickets {
			fmt.Printf("We only have %d tickets, so you can't book %v tickets", remainingTickets, remainingTickets)
		}

		remainingTickets = remainingTickets - userTicket
		//Array
		//bookings[0] = firstName + " " + lastname

		//Slice
		bookings = append(bookings, firstName+" "+lastname)

		fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastname, userTicket, email)

		fmt.Printf("%d tickets reamining for %s\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		//fmt.Printf("These are all our bookings: %v\n", bookings)

		fmt.Printf("The first names of bookings are: %s\n", firstNames)
	}
}
