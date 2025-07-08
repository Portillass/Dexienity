package main

import (

	"fmt"
	"time"
)
func main() {
	currentTime := time.Now()

	fmt.Println("Current Time:", currentTime)

	fmt.Println("Formated Date:", currentTime.Format("2006-01-02"))
	fmt.Println("Formated Time:", currentTime.Format("15:04:05"))
	fmt.Println("Dare and Time:" ,currentTime.Format("2006-01-02 15:04:05"))

}