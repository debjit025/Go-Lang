package main

import (
	"fmt"

	"examples.com/bank/fileops"
)

const accountBalanceText = "Balance.txt" //Accepting text file name

func main() {
	var choice int32
	var balance, err = fileops.GetFloatFromFile(accountBalanceText)
	var withdrawAmount, deposit float64
	if err != nil {
		err := err.Error()
		println("ERROR FOUND")
		println(err)
		println("----------------------------------")
		//panic("Cannot continue. Sorry!") //Make it seem more like an error and it will force the program to crash
	}
	fmt.Println("Welcome To BANK!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. EXIT")
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %.2f\n", balance)
		case 2:
			fmt.Print("Your Deposit: ")
			fmt.Scan(&deposit)
			if deposit < 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			balance += deposit
			fmt.Printf("Balance updated! New Balance: %.2f\n", balance)
			fileops.WriteFloatToFile(accountBalanceText, balance)
		case 3:
			fmt.Printf("Withdrawal Amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > balance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			balance -= withdrawAmount
			fmt.Printf("Balance updated! New Balance: %.2f\n", balance)
			fileops.WriteFloatToFile(accountBalanceText, balance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			//break: NOT APPLICABLE AS IT WILL NOT BREAK OUT OF THE LOOP
			return //Thats why we use return
		}
	}
}
