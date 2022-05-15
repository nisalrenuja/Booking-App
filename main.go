package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	const remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v conference tickets and remaining %v conference tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets from here to attend")

}
