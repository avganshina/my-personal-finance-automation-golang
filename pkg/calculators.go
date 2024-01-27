package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)


type TodaysBudget struct {
	Investments		float64
}

// GetUserInput retrieves user input from the console.
func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// GetUserInputInt retrieves integer user input from the console.
func GetUserInputInt(prompt string) int {
	inputStr := GetUserInput(prompt)
	inputInt, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer.")
		return GetUserInputInt(prompt)
	}
	return inputInt
}

// GenerateBudget creates a budget based on user input and assumptions.
func GenerateBudget(currentSalary string, age int, retirementAge int, desiredRetirementFund string) (*TodaysBudget, error) {
	// Convert input to float64
	currentSalaryFloat, err := strconv.ParseFloat(currentSalary, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting current salary to float64: %v", err)
	}

	desiredRetirementFundFloat, err := strconv.ParseFloat(desiredRetirementFund, 64)
	if err != nil {
	 	return nil, fmt.Errorf("error converting desired retirement fund to float64: %v", err)
	}

	// Basic assumptions
	annualRaisePercentage := 0.03
	inflationPercentage := 0.02
	adjustedRaisePercentage:= annualRaisePercentage-inflationPercentage

	// Calculate budget components
	investmentsNeeded := calulateNeededInvestments(age, retirementAge, desiredRetirementFundFloat, adjustedRaisePercentage)
	if investmentsNeeded > currentSalaryFloat/12{
		fmt.Print("Needed investments per month exceed income. Increase your income.")
	}
	// Create and return the budget
	return &TodaysBudget{
		Investments:          investmentsNeeded,
	}, nil
}


// Calculates how much you need to invest based on your age, salary, and desired retirement fund
func calulateNeededInvestments(age int, retirementAge int, desiredRetirementFund float64, adjustedRaisePercentage float64) float64 {
	yearsLeft := float64(retirementAge-age)
	n:=yearsLeft*12
	interestRate:=0.07
	r := interestRate/(12*100)
	denominator:= ((math.Pow((1.0+r),n) - 1.0)/r)
	PMT:= desiredRetirementFund / denominator
	AdjustedPMT := PMT / math.Pow((1.0 + adjustedRaisePercentage), yearsLeft)
	return AdjustedPMT
}
