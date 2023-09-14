package main

import (
	"fmt"
	"strings"
)

func IsValindInput(booking Booking) bool {
	isValidName := len(booking.firstName) >= 2 && len(booking.lastName) >= 2
	isValidEmail := strings.Contains(booking.email, "@") && strings.Contains(booking.email, ".")
	isValidTickets := booking.numOftickets > 0
	if !isValidName {
		fmt.Println("Sorry, your name must be at least 2 characters long")
		return false
	}
	if !isValidEmail {
		fmt.Println("Sorry, your email address is not valid")
		return false
	}
	if !isValidTickets {
		fmt.Println("Sorry, you must purchase at least one ticket")
		return false
	}
	return true
}

