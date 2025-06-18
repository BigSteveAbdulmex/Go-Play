package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	// Earnings before tax (ebt)
	ebt := revenue - expenses

	// Earnings after tax (eat or profit)
	eat := ebt * (1 - taxRate/100)

	// Ratio of ebt and eat
	ratio := ebt / eat

	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Earnings after tax:", eat)
	fmt.Println("Ratio of earnings before and after tax:", ratio)

}
