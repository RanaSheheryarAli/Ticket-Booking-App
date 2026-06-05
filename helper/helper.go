package helper

import (
	"fmt"
	"strings"
)

func ValidateUserData(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {

	isvalidateName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidateEmail := strings.Contains(email, "@")
	ivalidateUserTickets := userTickets > 0 && userTickets <= remainingTickets

	if isvalidateEmail && isvalidateName && ivalidateUserTickets {
		return true
	} else {
		return false
	}
}

func Greet(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("<------------Wellcome to our %v booking application------------>\n", conferenceName)

	fmt.Printf("<------------we have total of %v tickets and %v are still availavle------->\n", conferenceTickets, remainingTickets)

}
