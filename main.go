package main

import "fmt"

func main() {
	conferenceName := "PBKK Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var userName string
	var userTickets int

	userName = "Karina"
	userTickets = 3
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}