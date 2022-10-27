package main

import "fmt"

var menuNumber int
var Pin int
var accountBalance int

func main() {
	accountBalance = 20000
	initialPin := &Pin
	*initialPin = 0000
	welcome()
}

func welcome() {
	fmt.Println("*******Welcome to Nacho Atm*********")
	menu()
}

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Println("\n")
		i++
	}
}

func changePin() {
	newLine(1)
	fmt.Println("Please Input your new pin ...")
	_, err := fmt.Scan(&Pin)
	if err != nil {
		fmt.Println("Your new Pin is set")
	} else {
		fmt.Println("Please input a four digit pin.")
	}
	menu()
}

func validatePin() bool {
	fmt.Println("Please type your pin.")
	var pin int
	_, err := fmt.Scan(&pin)
	if Pin != 0000 {
		if err == nil && pin == Pin {
			return true
		}
		return false
	} else {
		fmt.Println("Change your Pin, odogwu")
		return false
	}
}

func AccountBalance() {

	newLine(1)
	if validatePin() {
		fmt.Println("Your Current Account balance is:", accountBalance)
	} else {
		fmt.Println("Incorrect Pin. Please Change your Pin")
		menu()
	}
	menu()
}

func withDrawFunds() {
	var amount int
	fmt.Println("How much do you want to withdraw ?")
	fmt.Scan(&amount)
	if validatePin() && amount < accountBalance {
		var newAccountNumber *int
		newAccountNumber = &accountBalance
		*newAccountNumber = *newAccountNumber - amount
		menu()
	} else {
		fmt.Println("Insufficient Account Balance or Incorrect Pin.")
		withDrawFunds()
	}
}

func depositFunds() {
	var amount int
	fmt.Println("How much do you want to withdraw ?")
	fmt.Scan(&amount)
	var newAccountNumber *int
	newAccountNumber = &accountBalance
	if validatePin() {
		*newAccountNumber = *newAccountNumber + amount
		menu()
	} else {
		fmt.Println("Incorrect Pin.Please Change your Pin.")
		depositFunds()
	}
}

func exitProgram() {
	fmt.Println("Bye bye")
}

func menu() {
	newLine(1)
	fmt.Println("Select Operation:")
	fmt.Println("1. Change Pin \t\t 2. Account Balance")
	fmt.Println("3. Withdraw Funds \t 4. Deposit Funds")
	fmt.Println("0. Exit program \t")

	_, err := fmt.Scan(&menuNumber)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
		menu()
	}

	switch menuNumber {
	case 1:
		changePin()
	case 2:
		AccountBalance()
	case 3:
		withDrawFunds()
	case 4:
		depositFunds()
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid input")
	}
}
