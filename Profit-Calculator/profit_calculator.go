package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Enter revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Enter expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Enter tax rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, eat, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.1f\n", ebt)
	fmt.Printf("Earnings after tax: %.1f\n", eat)
	fmt.Printf("Ratio of earnings before and after tax: %.3f\n", ratio)
	storeCalculatedResultsToFile(ebt, eat, ratio)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	// Earnings before tax (ebt)
	ebt := revenue - expenses
	// Earnings after tax (eat or profit)
	eat := ebt * (1 - taxRate/100)
	// Ratio of ebt and eat
	ratio := ebt / eat
	return ebt, eat, ratio
}

func storeCalculatedResultsToFile(ebt, eat, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nEAT: %.1f\nRatio: %.3f\n", ebt, eat, ratio)
	os.WriteFile("profits.txt", []byte(results), 0644)
}
