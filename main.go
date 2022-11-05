package main

import (
	"fmt"
	"math"
	"os"
)

type Account struct {
	balance  float64
	pin      string
	currency byte
}

var menuNumber int

var account = Account{balance: 0, pin: "0000", currency: '$'}

func main() {
	welcome()
	menu()
}

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Println()
		i++
	}
}

func welcome() {
	fmt.Println("\n*******{ Welcome to the World Bank ATM app by 🍾 }*******")
}

func viewAccountBalance() {
	fmt.Printf("\nYour account balance: %s%g\n", string(account.currency), account.balance)

	menu()
}

func withdrawFunds() {
	var amount float64
	fmt.Println("\nEnter the amount you want to withdraw: ")
	_, err := fmt.Scanln(&amount)

	if err != nil {
		fmt.Println("Error: Please enter the amount you want to withdraw")
		var discard string
		fmt.Scanln(&discard)
		withdrawFunds()
	}

	if math.Signbit(amount) {
		fmt.Println("\nPlease enter a valid amount")
		withdrawFunds()
	}

	if account.balance < amount {
		newLine(1)
		fmt.Printf("Lol...you get mind. You do not have up to %s%g in your account!", string(account.currency), amount)
		newLine(1)
		withdrawFunds()
	}

	account.balance = account.balance - amount
	newLine(1)
	fmt.Printf("You have successfully withdrawn: %s%g\n", string(account.currency), amount)
	menu()
}

func depositFunds() {
	var amount float64
	fmt.Println("\nEnter the amount you want to deposit: ")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println("Error: Please enter the amount you want to deposit")
		var discard string
		fmt.Scanln(&discard)
		depositFunds()
	}

	if math.Signbit(amount) {
		fmt.Println("\nPlease enter a valid amount")
		depositFunds()
	}

	account.balance = account.balance + amount
	newLine(1)
	fmt.Printf("You have successfully deposited: %s%g\n", string(account.currency), amount)
	menu()
}

func changePin() {
	var pin string
	fmt.Println("\nEnter your new pin: ")
	_, err := fmt.Scan(&pin)
	if err != nil {
		fmt.Println("Error: Please enter your new pin")
		var discard string
		fmt.Scanln(&discard)
		changePin()
	}

	account.pin = pin
	newLine(1)
	fmt.Printf("You new pin is: %s \n", account.pin)

	menu()
}

func viewPin() {
	newLine(1)
	fmt.Printf("You pin is: %s \n", account.pin)

	menu()
}

func exitProgram() {
	fmt.Println("\nGoodbye, Go and make more money!\n")
	os.Exit(0)
}

func menu() {
	newLine(1)
	fmt.Println("Select Operation:")
	fmt.Println("1. View Account Balance")
	fmt.Println("2. Withdraw Funds")
	fmt.Println("3. Deposit Funds")
	fmt.Println("4. Change Pin")
	fmt.Println("5. View Current Pin")
	fmt.Println("0. Exit ")

	_, err := fmt.Scan(&menuNumber)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
		//menu()
	}

	switch menuNumber {
	case 1:
		viewAccountBalance()
	case 2:
		withdrawFunds()
	case 3:
		depositFunds()
	case 4:
		changePin()
	case 5:
		viewPin()
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid input")
	}
}
