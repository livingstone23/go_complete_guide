package main

import (
	"fmt"
)

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to GO Bank")
	//for i := 0; i < 3; i++ {
	//In this way we can use for loop to limit the number of tries
	for {

		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is: ", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				//return

			}

			accountBalance += depositAmount
			fmt.Println("Your account balance is: ", accountBalance)

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount!. Must be greater than 0.")
				//return

			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")

			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Your account balance is: ", accountBalance)
			}
		default:
			fmt.Println(" Goodbye!")
			fmt.Println("Thank you for banking with us!.")
			return
		}

	}
}

/*

for {

		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {

			fmt.Println("Your account balance is: ", accountBalance)

		} else if choice == 2 {

			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your account balance is: ", accountBalance)

		} else if choice == 3 {

			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount!. Must be greater than 0.")
				//return
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Your account balance is: ", accountBalance)
			}

		} else {

			fmt.Println(" Goodbye!")
			//return
			break
		}
		/*
	}

	fmt.Println("Thank you for banking with us!.")

}

/*
//Untill here working with functions
func main() {
	fmt.Println("Hello, profit_Calculator!")

	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	ebt, profit, ratio := CalculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func CalculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
*/
