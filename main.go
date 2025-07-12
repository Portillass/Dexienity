package main

import (

	"fmt"  // Importing the fmt package for formatted I/O
	"time"  // Importing the time package to work with date and time
)
func main() {
	// Get the current local time
	currentTime := time.Now()
	
	// Print the full current time and date
	fmt.Println("Current Time:", currentTime)

	fmt.Println("Formated Date:", currentTime.Format("2006-01-02"))
	fmt.Println("Formated Time:", currentTime.Format("15:04:05"))
	fmt.Println("Dare and Time:" ,currentTime.Format("2006-01-02 15:04:05"))

}
