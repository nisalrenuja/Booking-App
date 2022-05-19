package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	var conferenceTickets int = 50
	var remainingTickets int = 50
	bookings := []string{}

	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v conference tickets and remaining %v conference tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets from here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("\nEnter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("\nEnter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("\nEnter your Email Address:")
		fmt.Scan(&email)

		fmt.Println("\nEnter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("First Names of bookings are: %v\n", firstNames)
	}
}
