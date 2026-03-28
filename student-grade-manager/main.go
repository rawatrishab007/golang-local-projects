package main

import (
	"fmt"
)

// storing data taken from the user
// struct for name and marks
type student struct {
	Name   string
	grades float64
}

func main() {
	//slice for storing data (act as a database )
	var students []student
	for { //making an infinite for loop so that we can exit the program on our own
		fmt.Println("\n--- Student Grade Manager ---")
		fmt.Println("1. Add Student")
		fmt.Println("2. View All Marks")
		fmt.Println("3. Delete a student")
		fmt.Println("4. Exit")
		fmt.Print("Select an option: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var name string
			var marks float64
			fmt.Println("Enter new student name :")
			fmt.Scan(&name)
			fmt.Println("Enter student grades")
			fmt.Scan(&marks)
			newStudent := student{Name: name, grades: marks}
			students = append(students, newStudent)
			fmt.Println("Student added sucessfully!")
		case 2:

			fmt.Println("----student list----")
			if len(students) == 0 {
				fmt.Println("No record found")
			} else {
				for _, s := range students {
					fmt.Printf("Name :%s | Grades: %.2f \n", s.Name, s.grades)
				}
			}
		case 3:
			if len(students) == 0 {
				fmt.Println("Nothing to delete😛")
			} else {
				// Show the list WITH the index number 'i'
				fmt.Println("Select index to delete:")
				for i, s := range students {
					fmt.Printf("[%d] Name: %s | Grade: %.2f\n", i, s.Name, s.grades)
				}
			}
			fmt.Println("Enter the index of student to delete ")
			var index int
			fmt.Scan(&index)
			if index < 0 || index >= len(students) {
				fmt.Println("Invalid index try again")
			} else {
				students = append(students[:index], students[index+1:]...)
			}
			fmt.Println("Student deleted sucessfully")
		case 4:
			fmt.Println("Exiting......goodbye ")
			return

		default:
			fmt.Println("Invalid choice ! try again")
		}
	}

}
