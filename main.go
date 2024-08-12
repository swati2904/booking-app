package main

import "fmt" // input/ outside functionality

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings = [50]string{"user1", "user2", "user3"}
	// alternative syntax
	// var bookings [50]string

	//slice
	var bookings []string

	// To check the data type of variables
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Define data type explicitly
	var firstName string
	var lastName string
	var email string
	var phoneNumber string
	var userTickets uint

	// Instead of defining the hardcore value ask the user to enter
	// & pointer is used to wait for user to enter the value instead of directly printing the value.

	// fmt.Println(remainingTickets)  // will print the value
	// fmt.Println(&remainingTickets) // will print the memory of the variable stored

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your phone number:")
	fmt.Scan(&phoneNumber)

	fmt.Println("Enter no of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first array: %v\n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n", len(bookings))

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first slice: %v\n", bookings[0])
	// fmt.Printf("slice type: %T\n", bookings)
	// fmt.Printf("slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v or an SMS on your phone Number %v\n", firstName, lastName, userTickets, email, phoneNumber)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all the our booking: %v\n", bookings)
}
