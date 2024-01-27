package main

import (
	"fmt"
	"log"
	"github.com/avganshina/personal-finance-golang/pkg"
)

func main() {
	// Get user input
	currentSalary := pkg.GetUserInput("Enter your current salary: ")
	age := pkg.GetUserInputInt("Enter your age: ")
	retirementAge := pkg.GetUserInputInt("Enter your planned retirement age: ")
	desiredRetirementFund := pkg.GetUserInput("Enter your desired retirement fund: ")

	// Create a budget
	budget, err := pkg.GenerateBudget(currentSalary, age, retirementAge, desiredRetirementFund)
	if err != nil {
		log.Fatal(err)
	}

	// Display the budget
	fmt.Println("\nGenerated Budget:")
	fmt.Printf("Invest per month: $%.2f\n", budget.Investments)
}
