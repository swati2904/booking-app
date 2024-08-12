package main

import "fmt"

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	// To check the data type of variables
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Define data type explicitly
	var userName string
	var userTickets int

	// Ask user for their name
	userName = "Jonathan"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

}
