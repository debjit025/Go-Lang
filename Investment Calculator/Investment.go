package main

import (
	"fmt"
	"math"
)

func main() {
	inflationRate := 6.5
	var investmentAmount, rateOfReturn, years float64
	fmt.Printf("Enter the Invested Amount: ")
	fmt.Scanln(&investmentAmount)
	fmt.Printf("Enter the Rate of Return: ")
	fmt.Scanln(&rateOfReturn)
	fmt.Printf("Enter the number of years of Investment: ")
	fmt.Scanln(&years)
	var returnAmount float64 = investmentAmount * math.Pow((1+rateOfReturn/100), years)
	var inflatedAmount float64 = returnAmount / math.Pow(1+inflationRate/100, years)
	fmt.Printf("After %d years of investemnt you would get %d Rupees \n", int(years), int(returnAmount))
	fmt.Printf("As per inflation rate of %f%%, your actual amount becomes %d Rupees", inflationRate, int(inflatedAmount))
}
