package main

import (
	"fmt"
)

type Wallet struct {
	ID int
	Balance float64
}

func (b *Wallet) Deposit(amount float64){
	b.Balance += amount
}

func (b *Wallet) Withdraw(amount float64) error {
	if b.Balance > amount{
		b.Balance -= amount
		return nil
	}

	return fmt.Errorf("Недостаточно средстве")
}

func NewWallet() *Wallet{
	var id int = 1
	var amnt float64 = 100.00



	return &Wallet{		
		ID: id,
		Balance: amnt,
	}
}
func Swap(a, b *int) {
	temp := *b
	
	*a = *b
	*b = temp 
}