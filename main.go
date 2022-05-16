package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	const remainingTickets int = 50

	fmt.Printf("\nconferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v conference tickets and remaining %v conference tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets from here to attend")

	var userName string
	var userTickets int

	userTickets = 2
	userName = "Tom"

	fmt.Printf("%v booked %v tikets\n", userName, userTickets)

}
