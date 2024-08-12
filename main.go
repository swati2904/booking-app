package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // List of map

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()

		//validate users input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//prevent invalid number of tickets book
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTickets(userTickets, firstName, lastName, email)
			//call func print first name
			firstNames := getFirstName()
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
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid. \n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}

	//looping again to grab only one element at a time
	//extract firstname from the list
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// Define data type explicitly
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Instead of defining the hardcore value ask the user to enter
	// & pointer is used to wait for user to enter the value instead of directly printing the value.

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets -= userTickets

	// create a map for a user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("########################")
}
