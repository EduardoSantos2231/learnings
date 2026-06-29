package utils

import (
	"errors"
	"fmt"
)

type ErrorInsufficientFunds struct {
	Balance float64
	Amount  float64
}

var ErrInvalidInitialBalance = errors.New("invalid initial balance")
var ErrInvalidAmount = errors.New("invalid withdraw value")
var ErrInvalidOwner = errors.New("invalid owner's name")
var ErrInvalidDeposit = errors.New("invalid deposit value")

type bankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) (*bankAccount, error) {
	var err error

	if initialBalance < 0 {
		err = fmt.Errorf("[ERROR] %w\n", ErrInvalidInitialBalance)
	} else if len(owner) < 2 {
		err = fmt.Errorf("[ERROR] %w\n", ErrInvalidOwner)
	}

	return &bankAccount{
		owner:   owner,
		balance: initialBalance,
	}, err

}

func (e *ErrorInsufficientFunds) Error() string {
	return "Insufficient Funds, make sure to provide a valid value"
}

func (b *bankAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("[ERROR] %w", ErrInvalidDeposit)
	}
	b.balance += amount
	return nil
}
func (b *bankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > b.balance {
		return &ErrorInsufficientFunds{
			Balance: b.balance,
			Amount:  amount,
		}
	}
	b.balance -= amount
	return nil
}

func (b *bankAccount) Balance() float64 {
	return b.balance
}

func (b *bankAccount) Owner() string {
	return b.owner
}

func ParseFloat(n float64) bool {
	return n != 0
}
