package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	const remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have", conferenceTickets, "conference tickets and remaining", remainingTickets, "conference tickets")
	fmt.Println("Get your tickets from here to attend")

}
