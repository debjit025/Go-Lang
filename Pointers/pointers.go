package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)
	fmt.Println("Adult Years:", calculateAdultYears(&age))
}

func calculateAdultYears(age *int) int {
	return *age - 18
}
