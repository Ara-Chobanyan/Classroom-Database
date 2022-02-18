package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func StartMenu() (string, error) {
	fmt.Println("--------------------------------------")
	fmt.Println("Please pick a option")
	fmt.Println("1) Create a new class")
	fmt.Println("2) Output the classes roster")
	fmt.Println("3) Find a student by there name")
	fmt.Println("4) Find a student by there ID")
	fmt.Println("5) Add a new student to the class")
	fmt.Println("6) Remove a student from the class (Id required)")
	fmt.Println("7) Update a students grade (ID required)")
	fmt.Println("8) Exit")
	fmt.Println("--------------------------------------")

	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()

	userChoice := userInput.Text()
	choice := userChoice

	if choice == "1" || choice == "2" || choice == "3" || choice == "4" || choice == "5" || choice == "6" || choice == "7" || choice == "8" {
		return choice, nil
	} else {
		return "", errors.New("Invalid input")
	}
}
