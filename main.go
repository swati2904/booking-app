package main

import "fmt" // input/ outside functionality

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
	var firstName string
	var userTickets int

	// Instead of defining the hardcore value ask the user to enter
	// & pointer is used to wait for user to enter the value instead of directly printing the value.

	// fmt.Println(remainingTickets)  // will print the value
	// fmt.Println(&remainingTickets) // will print the memory of the variable stored

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	// Ask user for their name
	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", firstName, userTickets)

}
