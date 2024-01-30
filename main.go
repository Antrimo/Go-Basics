package main

import "fmt"

func main() {
	var name = "Go Conference"  //variable declaration
	const conferenceTicket = 50 //constant declaration
	var remainingTicket = 50

	fmt.Printf("name is %T, conferenceTicket is %T, remainingTicket is %T \n", name, conferenceTicket, remainingTicket)

	fmt.Printf("Welcome to %v booking application \n", conferenceTicket)
	fmt.Println("Total ticket available for", name, "is", conferenceTicket, "and remaining ticket is", remainingTicket)
	fmt.Println("Get your ticket now!!!")

	var userName string //ask user for their name
	var userTicket int  //ask user for number of ticket they want to buy

	userName = "John Doe" //assign user name to variable
	userTicket = 5        //assign user ticket to variable
	fmt.Printf("User %v booked %v ticket", userName, userTicket)

}
