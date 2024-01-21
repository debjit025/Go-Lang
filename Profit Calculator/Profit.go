package main

import (
	"fmt"
	"os"
)

//Goals
// 1. Validate User input
// 	=> Show error message and exit
// 		- No negative numbers
// 		-Not 0
// 2. Store calculated results in a file.

const fileName = "Earnings Calculation.txt"

func writeToFile(earningsBeforeTax, earningsAfterTax_profit, ratio float64) {
	// ebt := fmt.Sprint(earningsBeforeTax)
	// profit := fmt.Sprint(earningsAfterTax_profit)
	// rat := fmt.Sprint(ratio)
	// os.WriteFile(fileName, []byte("Earnings before Tax is: "+ebt+"\nProfit: "+profit+"\nRatio: "+rat), 0644)

	resultsText := fmt.Sprintf("Earnings before Tax is: %.1f\nProfit: %.1f\nRatio: %.3f", earningsBeforeTax, earningsAfterTax_profit, ratio)
	os.WriteFile(fileName, []byte(resultsText), 0644)
}

func validateUserInput(value float64) error {
	if value < 0 {
		panic("The Amount can't be negative!")
	}
	if value == 0 {
		panic("The Amount can't be 0!")
	}
	return nil
}

func userInput(inputText string) float64 {
	var inputVal float64
	fmt.Print(inputText)
	fmt.Scan(&inputVal)
	validateUserInput(inputVal)
	return inputVal
}

func financialCalculation(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax_profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax_profit
	return earningsBeforeTax, earningsAfterTax_profit, ratio
}

func main() {
	// var earningsBeforeTax, earningsAfterTax_profit, ratio float64
	// fmt.Printf("Enter your Revenue: ")
	// fmt.Scan(&revenue)
	revenue := userInput("Enter your Revenue: ")
	// fmt.Printf("Enter your Expenses: ")
	// fmt.Scan(&expenses)
	expenses := userInput("Enter your Expenses: ")
	// fmt.Printf(`Enter the Tax Rate: `)
	// fmt.Scan(&taxRate)
	taxRate := userInput("Enter the Tax Rate: ")

	earningsBeforeTax, earningsAfterTax_profit, ratio := financialCalculation(revenue, expenses, taxRate)

	fmt.Printf("Your Earnings Before Tax or EBT is: %.2f\n", earningsBeforeTax)
	fmt.Printf("Your Earnings After Tax or Profit is: %.2f\n", earningsAfterTax_profit)
	fmt.Printf("The Ratio is: %.2f\n", ratio)
	writeToFile(earningsBeforeTax, earningsAfterTax_profit, ratio)
}
