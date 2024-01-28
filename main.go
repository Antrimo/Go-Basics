package main

import "fmt"

func main() {
	var name = "Go Conference"  //variable declaration
	const conferenceTicket = 50 //constant declaration
	var remainingTicket = 50
	fmt.Printf("Welcome to %v booking application \n", conferenceTicket)
	fmt.Println("Total ticket available for", name, "is", conferenceTicket, "and remaining ticket is", remainingTicket)
	fmt.Println("Get your ticket now!!!")

}
