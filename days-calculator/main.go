package main

import (
	"fmt"

	"time"
)

// intialize the struct for storing the data
type user struct {
	birthYear int
	AgeInDays int
}

// make function for the calculation
func (u *user) calculateAGE() {
	currentyear := time.Now().Year() //take real time date and convert it into year
	totalDays := 0
	for year := u.birthYear; year < currentyear; year++ {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			totalDays += 366
		} else {
			totalDays += 365
		}

	}
	u.AgeInDays = totalDays //update the counted days

}

// main function to print the output
func main() {

	var yearInput int //telling the lang to take input
	fmt.Print("Enter your birth year in here  (e.g. 2000): ")
	fmt.Scan(&yearInput)                 //scan the terminal to take input
	person := user{birthYear: yearInput} //assigning the values
	person.calculateAGE()
	fmt.Printf("Calculation complete \n")
	fmt.Printf("Year of birth :%d\n", person.birthYear)
	fmt.Printf("Age in days :%d\n", person.AgeInDays)
}
