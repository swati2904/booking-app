package main

import (
	"fmt" // input/ outside functionality
	"strings"
)

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

	// for loop

	for {
		// Define data type explicitly
		var firstName string
		var lastName string
		var email string
		var phoneNumber string
		var userTickets uint

		// Instead of defining the hardcore value ask the user to enter
		// & pointer is used to wait for user to enter the value instead of directly printing the value.

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

		//validate users input

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidPhoneNumber := len(phoneNumber) > 10
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//prevent invalid number of tickets book
		if isValidName && isValidEmail && isValidPhoneNumber && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v or an SMS on your phone Number %v\n", firstName, lastName, userTickets, email, phoneNumber)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			//looping again to grab only one element at a time
			//extract firstname from the list
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("First name or last name you entered is too short.\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address you entered is incorrect. \n")
			}
			if !isValidPhoneNumber {
				fmt.Printf("Phone number you entered is inValid. \n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid. \n")
			}
		}
	}
}
