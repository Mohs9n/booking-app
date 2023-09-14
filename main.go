package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Booking struct {
	firstName string
	lastName string
	email string
	numOftickets uint
}

var bookings []Booking

var wg = sync.WaitGroup{}

func main() {
	var confName = "Go Conference"
	const confTickets uint = 50
	var remainingTickets = confTickets
	greetUsers(confName, remainingTickets)
	for remainingTickets > 0 {
		booking := getUserInput()
		if !IsValindInput(booking) {
			continue
		}
		remainingTickets = bookTickets(remainingTickets, booking, confName)
		wg.Add(1)
		go sendTicketEmail(booking)
		printFirstNames(bookings)
	}
	println("Sorry, there are no more tickets available")
	wg.Wait()
}

func greetUsers(confName string, remainingTickets uint) {
	fmt.Printf("Welcome to the %v booking application!\n", confName)
	fmt.Println("We have a total of", remainingTickets, "tickets remaining.")
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []Booking) {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Println("The following people have booked tickets:", strings.Join(firstNames, ", "))
}

func getUserInput() Booking {
	var booking Booking

	fmt.Println("Please enter your first name")
	fmt.Scanln(&booking.firstName)

	fmt.Println("Please enter your last name")
	fmt.Scanln(&booking.lastName)

	fmt.Println("Please enter your email address")
	fmt.Scanln(&booking.email)

	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scanln(&booking.numOftickets)

	return booking
}

func bookTickets(remainingTickets uint, booking Booking, confName string) uint {
	remainingTickets = remainingTickets - booking.numOftickets
	fmt.Printf("Thank you %v %v, you have purchased %v tickets for the %v conference. Please check your email %v for confirmation.\n", booking.firstName, booking.lastName, booking.numOftickets, confName, booking.email)
	fmt.Println("We have a total of", remainingTickets, "tickets remaining.")
	
	bookings = append(bookings, booking)
	fmt.Printf("list of bookings: %v\n", bookings)
	return remainingTickets
}

func sendTicketEmail(booking Booking) {
	var ticket = fmt.Sprintf("%v", booking)
	fmt.Println("Sending email...")
	time.Sleep(10 * time.Second)
	fmt.Printf("\nSent %v to %v\n", ticket, booking.email)
	fmt.Println("Email sent!")
	wg.Done()
}