package main

import (

	"fmt"  // Importing the fmt package for formatted I/O
	"time"  // Importing the time package to work with date and time
)
// main function
func main() {
	// Get the current local time
	currentTime := time.Now()
	
	// Print the full current time and date
	fmt.Println("Current Time:", currentTime)

	// Format and print the date in YYYY-MM-DD format
	fmt.Println("Formated Date:", currentTime.Format("2006-01-02"))
	
	// Format and print the time in HH:MM:SS (24-hour) format
	fmt.Println("Formated Time:", currentTime.Format("15:04:05"))

	// Format and print both date and time together
	fmt.Println("Dare and Time:" ,currentTime.Format("2006-01-02 15:04:05"))

}
