package main

import (
	"fmt"
	"strings"
)

func main() {

	var conference = "Go Conference"
	const tickets = 50
	var remaining_tickets uint = 50
	var bookings = []string{} // slice
	// var bookings [50]string  // array of fixed size

	fmt.Println("--------------------------------------------------------------------------")

	fmt.Println("                         welcome to", conference, "!                       ")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("Get your tickets here!")
	//fmt.Println("Number of tickets available:", tickets)
	fmt.Printf("Number of tickets available: %v\n", tickets)
	fmt.Println("Number of tickets remaining:", remaining_tickets)

	// Note

	/*
		we will be using %v for printing the value of the variable

		Now we will betaking data from user using fmt.Scan

		now we will be using loop to amke our code more efficient

	*/

	// for loop

	for remaining_tickets > 0 && len(bookings) < 50 {

		Firstname, lastname, email, usertickets := getUserInput()

		isvalid, isvalidd := validate(Firstname, lastname, email, usertickets)

		if isvalid && isvalidd {

			bookticket(remaining_tickets, bookings, Firstname, lastname, email, usertickets)

		} else if usertickets > remaining_tickets {
			fmt.Println("Sorry! We only have", remaining_tickets, "tickets left.")
			//break
			continue

			// it will alow the user to enter the data again

		} else if remaining_tickets == 0 {

			//end the program

			fmt.Println("Sorry! No more tickets available. Please try next year.")
			break

		} else {
			fmt.Println("Please enter the valid data")
			continue
		}

		//Note :-

		/*

			as our array is of fixed size we will be using slices for gettinga any amount of data

			gonna use function of string package for getting the first name of the user



			now we will be using if else statement for checking the condition

			noticketsremaining := remaining_tickets == 0

		*/

		getfirstname(bookings)

	}

}

func getfirstname(bookings []string) []string {

	firstnames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	fmt.Printf("they are all the bookings %v\n", firstnames)
	return firstnames

}

func validate(Firstname string, lastname string, email string, usertickets uint) (bool, bool) {

	isvalid := len(Firstname) >= 2 && len(lastname) >= 2 && strings.Contains(email, "@")
	isvalidd := strings.Contains(email, ".") && usertickets > 0

	return isvalid, isvalidd

}

func bookticket(remaining_tickets uint, bookings []string, Firstname string, lastname string, email string, usertickets uint) {

	remaining_tickets = remaining_tickets - usertickets
	bookings = append(bookings, Firstname+" "+lastname)

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Bookings: %v\n", bookings)
	fmt.Println("Number of tickets remaining:", remaining_tickets)
	fmt.Printf("Hello %v %v, you have booked %v tickets and you will get confirmation of tickets on your %v\n", Firstname, lastname, usertickets, email)

}

func getUserInput() (string, string, string, uint) {
	var Firstname string
	var usertickets uint
	var lastname string
	var email string

	fmt.Println("Enter your first name:")
	fmt.Scan(&Firstname)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book:")
	fmt.Scan(&usertickets)

	return Firstname, lastname, email, usertickets

	//Extra code :-

	/*
		we will be using %T for checking the type of the variable

		fmt.Printf("username is of type %T\n", username)

		bookings[0] = Firstname + " " + lastname

		checking the validity of the data entered by user

		//fmt.Printf("type of slice %T\n", bookings)

		fmt.Printf("first value is %v\n", bookings[0])

		//fmt.Printf("last name is  %v\n", bookings[1])

		//fmt.Printf("length of slice  is %v\n", len(bookings))

	*/
}
