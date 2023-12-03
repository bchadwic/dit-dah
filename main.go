package main

import "fmt"

func main() {
	var message string
	fmt.Print("Type a message to send ")
	fmt.Scan(&message)
	fmt.Println("Your message to be sent is:", message)
}
