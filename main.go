// code needs to be in package
package main

// import package to use print, printf, scan, etc.
import (
	"LearnGolang/helper"
	"fmt"
	"sync"
	"time"
)

// search for packages
// https://pkg.go.dev/

// store in a variable name
var conferenceName = "Go Conference"

const conferenceTickets int = 50

// these variables only accessible in main.go
var remainingTickets uint = 50

// creating an empty list of maps, need to define size but it is dynamic and can increase
//var bookings = make([]map[string]string, 0)
// creating a struct
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// entry point, begining of execution
func main() {

	// remainingTickets = -1
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers()

	for {
		// get user input funciton call here
		firstName, lastName, email, userTickets := getUserInput()

		// validateUserInput call function here
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			//Example Concurrency
			//weight group
			wg.Add(1)

			go sendTicket(userTickets, firstName, lastName, email)
			//fmt.Printf("This is in the bookings array: %v\n", bookings)

			// call function printFirstNames
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program tickets sold out
				fmt.Printf("Our conference is booked out. Come back next time.\n")
				// terminate for loop
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}
			//continue to the next iteration and skips the code below, I moved order of code so no longer needed, note for future
			// continue
		}
	}
	// waitgroup till sendTicket is executed
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// get firstname value with a map
		//firstNames = append(firstNames, booking["firstName"])
		// get firstname value with a struct
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	// pass the reference, assign userName to memory address
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// need same type for operation
	remainingTickets = remainingTickets - userTickets

	//create a map for a user- empty
	// var userData = make(map[string]string)
	//change to a struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// define and assign key values when you have a map
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	// need to convert our uint tickets to a string
	// 10 is for base 10 - base numbers
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// simulate loading sending email
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address: %v\n", ticket, email)
	fmt.Println("################")
	// tells weightgroup it finished, removes the wg.Add(1) so it no longer weights, decreasing counter Add(1) to Add(0)
	wg.Done()
}
