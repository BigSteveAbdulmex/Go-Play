package main

import "fmt"

func main() {
	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	ebt, eat, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.1f\n", ebt)
	fmt.Printf("Earnings after tax: %.1f\n", eat)
	fmt.Printf("Ratio of earnings before and after tax: %.3f\n", ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
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
