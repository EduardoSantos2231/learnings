package main

import (
	"02-bank-account/utils"
	"errors"
	"flag"
	"fmt"
	"os"
)

type flags struct {
	owner    string
	deposit  float64
	withdraw float64
}

func main() {
	options := collectFlags()
	account, err := utils.NewBankAccount(options.owner, 5000)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidOwner) {
			fmt.Printf("\n%v\n", err)
			os.Exit(2)
			return
		}
	}
	if utils.ParseFloat(options.deposit) {
		err := account.Deposit(options.deposit)
		if errors.Is(err, utils.ErrInvalidDeposit) {
			fmt.Printf("%v\n", err)
			os.Exit(2)
			return
		}
	}

	if utils.ParseFloat(options.withdraw) {
		err := account.Withdraw(options.withdraw)
		var insuff *utils.ErrorInsufficientFunds
		if errors.As(err, &insuff) {
			fmt.Printf("Balance: %.1f\nAttempted: %.2f\n", insuff.Balance, insuff.Amount)
			os.Exit(1)
			return
		}
		if errors.Is(err, utils.ErrInvalidAmount) {
			fmt.Printf("%v\n", err)
			os.Exit(2)
		}
	}

	fmt.Printf("\nDone\nBalance: %.2f\n", account.Balance())
}

func collectFlags() *flags {
	var options flags
	flag.StringVar(&options.owner, "owner", "", "specify the name of the account's owner")
	flag.Float64Var(&options.deposit, "deposit", 0.0, "specify the deposit value")
	flag.Float64Var(&options.withdraw, "withdraw", 0.0, "specify the withdraw value")
	flag.Parse()
	return &options
}
