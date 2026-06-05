package main

import (
	"Booking_App/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)

var firstNamesToShow []string
var err error

var bookings []UserData

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var Wg = sync.WaitGroup{}

func main() {

	helper.Greet(conferenceName, conferenceTickets, remainingTickets)

	// for {

	firstName, lastName, email, userTickets := getUserData()
	if helper.ValidateUserData(firstName, lastName, email, userTickets, remainingTickets) == true {

		bookTickets(firstName, lastName, userTickets, email)
		Wg.Add(1)
		go sendingTickets(firstName, lastName, email, userTickets)

		printAllUserFirstName()

		if remainingTickets == 0 {
			fmt.Println("All the tickets are sold out thank you!")
			fmt.Printf("All user are %v", bookings)
			// break
		}
	} else {
		fmt.Println("Invalide User data")

	}
	// }
	Wg.Wait()
}

func printAllUserFirstName() {

	firstNamesToShow := []string{}

	for _, indexvalue := range bookings {
		firstNamesToShow = append(firstNamesToShow, indexvalue.firstName)
	}
	fmt.Printf("All user %v \n", firstNamesToShow)

}

func getUserData() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter how many tickets you want to buy")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v you buy %v tickets. you will get confirmation email at %v shortly\n", firstName, lastName, userTickets, email)

	fmt.Printf("Remaining tickets are %v \n", remainingTickets)
	fmt.Println("")

}

func sendingTickets(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	fmt.Println("<---------- Sending Tickets ---------->")
	fmt.Println(" ")
	var ticket = fmt.Sprintf("Congratulations for bought %v tickects for our todays events: %v %v \n", userTickets, firstName, lastName)
	fmt.Printf("Ticket Sent: %v \n", ticket)

	fmt.Printf("<---------- %v ----------> \n", email)
	Wg.Done()

}
